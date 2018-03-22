package repositories

import (
	"github.com/alex19861108/burj/burj_center/iris/common/database"
	"github.com/alex19861108/burj/burj_center/iris/errors"
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"gopkg.in/mgo.v2/bson"
)

type JobRepository interface {
	FindAll(m bson.M) (results []proto.Job)
	FindByID(id string) (job proto.Job, err error)
	InsertOrUpdate(job proto.Job) (updatedJob proto.Job, err error)
	DeleteByID(id string) (err error)
}

func NewJobRepository() JobRepository {
	return &jobMongoRepository{ColJob: "job"}
}

type jobMongoRepository struct {
	ColJob string
}

func (r *jobMongoRepository) FindAll(m bson.M) (results []proto.Job) {
	session, err := database.NewSession()
	if err != nil {
		return
	}
	defer session.Close()

	err = session.C(r.ColJob).Find(m).Sort("-_id").All(&results)
	if err != nil {
		panic(err)
	}
	return
}

func (r *jobMongoRepository) FindByID(id string) (job proto.Job, err error) {
	session, err := database.NewSession()
	if err != nil {
		return
	}
	defer session.Close()

	query := session.C(r.ColJob).FindId(id)
	if n, _ := query.Count(); n > 0 {
		query.One(&job)
		return job, nil
	} else {
		return proto.Job{}, errors.ERROR_DB_NOT_FOUND_NODE
	}
}

func (r *jobMongoRepository) InsertOrUpdate(job proto.Job) (proto.Job, error) {
	session, err := database.NewSession()
	if err != nil {
		return job, err
	}
	defer session.Close()

	jobs := r.FindAll(bson.M{"_id": job.Id})
	if len(jobs) == 0 {
		job.Id = bson.NewObjectId().Hex()
		err := session.C(r.ColJob).Insert(job)
		return job, err
	}

	current := jobs[0]
	err = session.C(r.ColJob).Update(bson.M{"_id": current.Id}, job)
	return current, err
}

func (r *jobMongoRepository) DeleteByID(id string) error {
	session, err := database.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.C(r.ColJob).Remove(bson.M{"_id": id})
	return err
}
