package lib

import (
	"bytes"
	"compress/bzip2"
	"github.com/Juminiy/kube/pkg/util"
	bz2 "github.com/dsnet/compress/bzip2"
	"io"
	"testing"
)

func TestLZ4(t *testing.T) {

}

var paperText = `
2 Data Models & Query Languages
For our discussion here, we group the research and development thrusts in data models and query languages
for database into eight categories.
2.1 MapReduce Systems
Google constructed their MapReduce (MR) framework
in 2003 as a “point solution” for processing its periodic
crawl of the internet [122]. At the time, Google had
little expertise in DBMS technology, and they built MR
to meet their crawl needs. In database terms, Map is a
user-defined function (UDF) that performs computation
and/or filtering while Reduce is a GROUP BY operation.
To a first approximation, MR runs a single query:
SELECT map() FROM crawl_table GROUP BY reduce()
Google’s MR approach did not prescribe a specific
data model or query language. Rather, it was up to the
Map and Reduce functions written in a procedural MR
program to parse and decipher the contents of data files.
There was a lot of interest in MR-based systems at
other companies in the late 2000s. Yahoo! developed
an open-source version of MR in 2005, called Hadoop.
It ran on top of a distributed file system HDFS that was
a clone of the Google File System [134]. Several startups were formed to support Hadoop in the commercial
marketplace. We will use MR to refer to the Google
implementation and Hadoop to refer to the open-source
version. They are functionally similar.
There was a controversy about the value of Hadoop
compared to RDBMSs designed for OLAP workloads.
This culminated in a 2009 study that showed that data
warehouse DBMSs outperformed Hadoop [172]. This
generated dueling articles from Google and the DBMS
community [123, 190]. Google argued that with careful engineering, a MR system will beat DBMSs, and a
user does not have to load data with a schema before
running queries on it. Thus, MR is better for “one shot”
tasks, such as text processing and ETL operations. The
DBMS community argued that MR incurs performance
problems due to its design that existing parallel DBMSs
already solved. Furthermore, the use of higher-level
languages (SQL) operating over partitioned tables has
proven to be a good programming model [127].
A lot of the discussion in the two papers was on implementation issues (e.g., indexing, parsing, push vs. pull
query processing, failure recovery). From reading both
papers a reasonable conclusion would be that there is a
place for both kinds of systems. However, two changes
in the technology world rendered the debate moot.
The first event was that the Hadoop technology and
services market cratered in the 2010s. Many enterprises
spent a lot of money on Hadoop clusters, only to find
there was little interest in this functionality. Developers
found it difficult to shoehorn their application into the
restricted MR/Hadoop paradigm. There were considerable efforts to provide a SQL and RM interface on top
of Hadoop, most notable was Meta’s Hive [30, 197].
The next event occurred eight months after the CACM
article when Google announced that they were moving
their crawl processing from MR to BigTable [164]. The
reason was that Google needed to interactively update
its crawl database in real time but MR was a batch system. Google finally announced in 2014 that MR had no
place in their technology stack and killed it off [194].
The first event left the three leading Hadoop vendors
(Cloudera, Hortonworks, MapR) without a viable product to sell. Cloudera rebranded Hadoop to mean the
whole stack (application, Hadoop, HDFS). In a further
sleight-of-hand, Cloudera built a RDBMS, Impala [150],
on top of HDFS but not using Hadoop. They realized
that Hadoop had no place as an internal interface in a
SQL DBMS, and they configured it out of their stack
with software built directly on HDFS. In a similar vein,
MapR built Drill [22] directly on HDFS, and Meta created Presto [185] to replace Hive.
Discussion: MR’s deficiencies were so significant that
it could not be saved despite the adoption and enthusiasm from the developer community. Hadoop died
about a decade ago, leaving a legacy of HDFS clusters
in enterprises and a collection of companies dedicated
to making money from them. At present, HDFS has
lost its luster, as enterprises realize that there are better
distributed storage alternatives [124]. Meanwhile, distributed RDBMSs are thriving, especially in the cloud.
Some aspects of MR system implementations related
to scalability, elasticity, and fault tolerance are carried
over into distributed RDBMSs. MR also brought about
the revival of shared-disk architectures with disaggregated storage, subsequently giving rise to open-source
file formats and data lakes (see Sec. 3.3). Hadoop’s limitations opened the door for other data processing platforms, namely Spark [201] and Flink [109]. Both systems started as better implementations of MR with procedural APIs but have since added support for SQL [105].
2.2 Key/Value Stores
The key/value (KV) data model is the simplest model
possible. It represents the following binary relation:
(key,value)
A KV DBMS represents a collection of data as an associative array that maps a key to a value. The value is
typically an untyped array of bytes (i.e., a blob), and the
DBMS is unaware of its contents. It is up to the application to maintain the schema and parse the value into
its corresponding parts. Most KV DBMSs only provide
get/set/delete operations on a single value.
In the 2000s, several new Internet companies built
their own shared-nothing, distributed KV stores for nar
`
var srcBytes = []byte(util.MagicStr)
var dstBytes = []byte(``)

// Error on MacOS
func TestBZip2Decode(t *testing.T) {
	rd := bzip2.NewReader(bytes.NewReader(dstBytes))
	rdN, err := rd.Read(dstBytes)
	if err != nil {
		t.Error(err)
	}
	t.Logf("read %d Bytes, encoded Info: %s", rdN, dstBytes)
}

func TestUnofficialBZip2Encode(t *testing.T) {
	strBuf := bytes.NewBuffer(dstBytes)
	wtr, err := bz2.NewWriter(strBuf, &bz2.WriterConfig{Level: bz2.BestCompression})
	if err != nil {
		t.Error(err)
	}
	wtN, err := wtr.Write(srcBytes)
	if err != nil {
		t.Error(err)
	}
	util.SilentCloseIO("bz2 writer", wtr)
	t.Logf("%dBytes -> %dBytes, encoded To:\n%s", wtN, len(strBuf.Bytes()), strBuf.Bytes())
}

func TestUnofficialBZip2Decode(t *testing.T) {
	bz2Rdr, err := bz2.NewReader(bytes.NewReader(dstBytes), &bz2.ReaderConfig{})
	if err != nil {
		t.Error(err)
	}
	strBuf := bytes.NewBuffer(dstBytes)
	cpN, err := io.Copy(strBuf, bz2Rdr)
	if err != nil {
		t.Error(err)
	}
	util.SilentCloseIO("bz2 reader", bz2Rdr)
	t.Logf("read %d Bytes, decoded To: %s", cpN, strBuf.String())
}
