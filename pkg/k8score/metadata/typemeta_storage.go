package metadata

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func TypeMetaStorageClass() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "StorageClass",
		APIVersion: "storage.k8s.io/v1",
	}
}

func TypeMetaVolumeAttachment() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "VolumeAttachment",
		APIVersion: "storage.k8s.io/v1",
	}
}

func TypeMetaCSINode() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "CSINode",
		APIVersion: "storage.k8s.io/v1",
	}
}

func TypeMetaCSIDriver() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "CSIDriver",
		APIVersion: "storage.k8s.io/v1",
	}
}

func TypeMetaCSIStorageCapacity() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "CSIStorageCapacity",
		APIVersion: "storage.k8s.io/v1",
	}
}
