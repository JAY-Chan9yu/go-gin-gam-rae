package cosmetic

import (
	"context"
	"fmt"
	infra "github.com/JAY-Chan9yu/go-gin-gam-rae/infrastructs"
	pb "github.com/JAY-Chan9yu/go-gin-gam-rae/proto/cosmetic"
	"log"
	"strconv"
)

type Cosmetic struct {
	pb.UnimplementedCosmeticServiceServer
}

type CosmeticDto struct {
	id          string
	name        string
	description string
	price       int32
}

func (s *Cosmetic) DeleteCosmetic(ctx context.Context, in *pb.DeleteCosmeticRequest) (*pb.DeleteCosmeticReply, error) {
	db := infra.GetDBConnection()
	defer db.Close()

	result, err := db.Exec("delete FROM Cosmetic WHERE id = " + "'" + in.Uuid + "'")

	if err != nil {
		log.Fatal(err)
	}

	nRow, err := result.RowsAffected()
	fmt.Println(nRow)
	t := strconv.Itoa(int(nRow))
	return &pb.DeleteCosmeticReply{Message: "delete count: " + t}, nil
}

func (s *Cosmetic) UpdateCosmetic(ctx context.Context, in *pb.UpdateCosmeticRequest) (*pb.UpdateCosmeticReply, error) {
	db := infra.GetDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Cosmetic SET name=?, description=?, price=? WHERE id=?")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.Name, in.Description, in.Price, in.Uuid)
	if err != nil {
		log.Fatalf("failed to execute statement: %v", err)
	}

	nRow, err := result.RowsAffected()
	fmt.Println(nRow)
	return &pb.UpdateCosmeticReply{Uuid: in.Uuid, Name: in.Name, Description: in.Description, Price: in.Price}, nil
}

func (s *Cosmetic) CreateCosmetic(ctx context.Context, in *pb.CreateCosmeticRequest) (*pb.CreateCosmeticResponse, error) {
	db := infra.GetDBConnection()
	defer db.Close()

	result, err := db.Exec("INSERT INTO Cosmetic(name, description,price) VALUES(?, ?, ?)", in.Name, in.Description, in.Price)

	if err != nil {
		log.Fatal(err)
	}
	_, err = result.RowsAffected()
	return &pb.CreateCosmeticResponse{Message: ""}, nil
}

func (s *Cosmetic) ListCosmetics(ctx context.Context, in *pb.ListCosmeticsRequest) (*pb.ListCosmeticsResponse, error) {
	db := infra.GetDBConnection()
	defer db.Close()

	query := "SELECT COUNT(*) FROM cosmetic"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var cnt int32
	fmt.Println(rows)
	for rows.Next() {
		if err := rows.Scan(&cnt); err != nil {
			log.Fatal(err)
		}
	}

	query = "SELECT * FROM cosmetic"
	rows, err = db.Query(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
	cosmetics := make([]*pb.Cosmetics, cnt)

	i := 0
	for rows.Next() {
		var cosmeticDto CosmeticDto
		if err := rows.Scan(&cosmeticDto.id, &cosmeticDto.name, &cosmeticDto.description, &cosmeticDto.price); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name is %s\n", cosmeticDto.id, cosmeticDto.name)
		cosmetics[i] = &pb.Cosmetics{Id: cosmeticDto.id, Name: cosmeticDto.name, Description: cosmeticDto.description, Price: cosmeticDto.price}
		i += 1
	}

	return &pb.ListCosmeticsResponse{Data: cosmetics}, nil
}
