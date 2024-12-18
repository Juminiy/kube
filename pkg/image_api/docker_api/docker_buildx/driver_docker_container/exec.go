package driver_docker_container

import (
	"bufio"
	"bytes"
	"context"
	dockerapiprovider "github.com/Juminiy/kube/pkg/image_api/docker_api/api_provider"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	dockercli "github.com/docker/docker/client"
	dockerpkgstdcopy "github.com/docker/docker/pkg/stdcopy"
	"github.com/pkg/errors"
	"net"
)

type containerExec struct {
	dockerapiprovider.APIProvider
	ctx      context.Context
	idOrName string // container id or name

	// exec stdout
	stdout *bytes.Buffer
	// exec stderr
	stderr *bytes.Buffer
}

func (e *containerExec) RunBuild() error {
	return nil
}

func (e *containerExec) Bootstrap() error {
	crt, err := e.inspect()
	if err != nil {
		return err
	}
	if crtState := crt.State; crtState != nil {
		//"created", "running", "paused", "restarting", "removing", "exited", or "dead"
		if crtState.Running || crtState.Restarting || util.ElemIn(crtState.Status, "running", "restarting") {
			return nil
		} //else if crtState.Paused || crtState.Dead || crtState.ExitCode != 0 || util.ElemIn(crtState.Status, "paused", "exited", "dead") { }
	}

	return e.start()
}

var ErrContainerExecExitCode = errors.New("container exec exit code not 0")

type ExecRunResp struct {
	ExecCreateAttachResp
	container.ExecInspect
}

func (e *containerExec) run(cmd ...string) (resp ExecRunResp, err error) {
	execCreateAttach, err := e.exec(cmd...)
	if err != nil {
		return
	}
	defer util.SilentCloseIO("docker container exec netConn", execCreateAttach.Conn)
	_, err = dockerpkgstdcopy.StdCopy(e.stdout, e.stderr, execCreateAttach.Conn)
	if err != nil {
		err = errors.Wrap(err, "copy from exec attach net.Conn")
		return
	}

	resp.ExecCreateAttachResp = execCreateAttach
	execInspect, err := e.SDK.ContainerExecInspect(e.ctx, execCreateAttach.ExecID)
	if err != nil {
		return
	} else if execInspect.ExitCode != 0 {
		err = errors.Wrapf(ErrContainerExecExitCode, "exit code is: %d", execInspect.ExitCode)
		return
	}
	return
}

func (e *containerExec) inspect() (crt types.ContainerJSON, err error) {
	crt, err = e.SDK.ContainerInspect(e.ctx, e.idOrName)
	if dockercli.IsErrNotFound(err) {
		return crt, errors.Wrapf(err, "buildx-instance container[IdOrName:%s] not found", e.idOrName)
	}
	e.idOrName = crt.ID
	return
}

func (e *containerExec) start() error {
	return e.SDK.ContainerStart(e.ctx, e.idOrName, container.StartOptions{})
}

var ErrContainerExecCreateRespID = errors.New("container exec create response id empty")

type ExecCreateAttachResp struct {
	ExecID string
	Conn   net.Conn
	Reader *bufio.Reader
}

func (e *containerExec) exec(cmd ...string) (resp ExecCreateAttachResp, err error) {
	execIDResp, err := e.SDK.ContainerExecCreate(e.ctx, e.idOrName, container.ExecOptions{
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          cmd,
	})
	if err != nil {
		return
	} else if len(execIDResp.ID) == 0 {
		err = ErrContainerExecCreateRespID
		return
	}

	resp.ExecID = execIDResp.ID
	execAttach, err := e.SDK.ContainerExecAttach(e.ctx, execIDResp.ID, container.ExecAttachOptions{})
	if err != nil {
		return
	}

	resp.Reader = execAttach.Reader
	resp.Conn = execAttach.Conn
	return
}
