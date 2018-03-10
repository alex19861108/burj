package repositories

import (
	"github.com/alex19861108/burj/burj_center/iris/common/database"
	"github.com/alex19861108/burj/burj_center/iris/errors"
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"gopkg.in/mgo.v2/bson"
)

type NodeRepository interface {
	FindAll(m bson.M) (results []*proto.Node, err error)
	FindByID(id string) (node proto.Node, err error)
	InsertOrUpdate(node proto.Node) (updatedNode proto.Node, err error)
	DeleteByID(id bson.ObjectId) (err error)
	DeleteBy(m bson.M) error
	DropCollection() error
}

// NewNodeRepository returns a new node memory-based repository,
// the one and only repository type in our example.
func NewNodeRepository() NodeRepository {
	return &nodeMongoRepository{ColNode: "node"}
}

// nodeMongoRepository is a "NodeRepository"
// which manages the nodes using the memory data source (map).
type nodeMongoRepository struct {
	ColNode string
}

func (r *nodeMongoRepository) FindAll(m bson.M) (results []*proto.Node, err error) {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return
	}
	defer session.Close()

	err = session.C(r.ColNode).Find(m).All(&results)
	if err != nil {
		panic(err)
	}
	return
}

func (r *nodeMongoRepository) FindByID(id string) (node proto.Node, err error) {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return
	}
	defer session.Close()

	query := session.C(r.ColNode).FindId(id)
	if n, _ := query.Count(); n > 0 {
		query.One(&node)
		return node, nil
	} else {
		return proto.Node{}, errors.ERROR_DB_NOT_FOUND_NODE
	}
}

func (r *nodeMongoRepository) InsertOrUpdate(node proto.Node) (proto.Node, error) {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return node, err
	}
	defer session.Close()

	nodes, err := r.FindAll(bson.M{"addr.http": node.Addr.Http})
	if len(nodes) == 0 || err != nil {
		node.Id = bson.NewObjectId().Hex()
		err := session.C(r.ColNode).Insert(node)
		return node, err
	}

	node.Id = nodes[0].Id
	err = session.C(r.ColNode).Update(bson.M{"addr.http": node.Addr.Http}, node)
	if err != nil {
		panic(err)
	}
	return node, err
}

func (r *nodeMongoRepository) DeleteByID(id bson.ObjectId) error {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return err
	}
	defer session.Close()

	err = session.C(r.ColNode).Remove(bson.M{"_id": id})
	return err
}

func (r *nodeMongoRepository) DeleteBy(m bson.M) error {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return err
	}
	defer session.Close()

	err = session.C(r.ColNode).Remove(m)
	return err
}

func (r *nodeMongoRepository) DropCollection() error {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return err
	}
	defer session.Close()

	err = session.C(r.ColNode).DropCollection()
	return err
}
