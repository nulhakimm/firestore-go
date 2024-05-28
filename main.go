package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Use your service account file
	sa := option.WithCredentialsFile("./serviceAccountKey.json")

	client, err := firestore.NewClient(ctx, "latihan-2-89ee6", sa)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// docData := map[string]interface{}{
	// 	"name": "Null Hakim",
	// 	"age":  22,
	// }
	// err = AddDocument(ctx, client, "users", docData)
	// if err != nil {
	// 	log.Fatalf("Failed to add document: %v", err)
	// }

	err = getDocument(ctx, client, "users", "0UQZzbzNNcoSsOHkDhpX")
	if err != nil {
		log.Fatalf("Users Not Found : %v", err)
	}

}

func AddDocument(ctx context.Context, client *firestore.Client, collectionName string, docData map[string]interface{}) error {
	_, _, err := client.Collection(collectionName).Add(ctx, docData)
	if err != nil {
		return fmt.Errorf("failed adding document: %v", err)
	}
	fmt.Println("Document successfully written")
	return nil
}

func getDocument(ctx context.Context, client *firestore.Client, collectionName, docID string) error {
	doc, err := client.Collection(collectionName).Doc(docID).Get(ctx)
	if err != nil {
		return fmt.Errorf("failed getting document: %v", err)
	}
	fmt.Printf("Document data: %#v\n", doc.Data())
	return nil
}
