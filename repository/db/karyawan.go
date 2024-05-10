package db

import (
	"api/config"
	"api/model"
	"errors"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDataKaryawan(requestData model.DataKaryawan) error {
	db := mongo.MongoConnect(DBATS)
	insertedID := mongo.InsertOneDoc(db, config.KaryawanColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func CountDataKaryawan(version string) error {
	db := mongo.MongoConnect(DBATS)
	filter := bson.M{"codename": version}
	checkData, err := mongo.CountDocuments(db, config.KaryawanColl, filter)
	if err != nil {
		return err // Return the error if there's an issue with counting documents
	}
	if checkData > 0 {
		return errors.New("data already exists") // Return an error if data already exists
	}
	return nil
}

func GetDataKaryawanFilter(filter bson.M) ([]model.DataKaryawan, error) {
	db := mongo.MongoConnect(DBATS)
	var datas []model.DataKaryawan
	err := mongo.GetAllDocByFilter[model.DataKaryawan](db, config.KaryawanColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneDataKaryawanFilter(filter bson.M) (model.DataKaryawan, error) {
	db := mongo.MongoConnect(DBATS)
	var data model.DataKaryawan
	err := mongo.GetOneDocByFilter[model.DataKaryawan](db, config.KaryawanColl, filter, &data)
	if err != nil {
		return model.DataKaryawan{}, err
	}
	return data, nil
}
