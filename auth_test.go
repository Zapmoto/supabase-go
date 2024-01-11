package supabase

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)


func TestDeleteUser(t *testing.T) {
	t.Skip()
	godotenv.Load(".env")
	url := os.Getenv("SUPABASE_URL")
	serviceRole := os.Getenv("SUPABASE_SERVICE_ROLE")
	client := CreateClient(url, serviceRole)

	if err := client.Auth.DeleteUser(context.Background(), "baa1d7c8-571a-4b8e-8be5-61f3e9bf8691"); err != nil {
		t.Fatal(err)
	}
}

func TestGetUserFromEmail(t *testing.T) {
	t.Skip()
	godotenv.Load(".env")
	url := os.Getenv("SUPABASE_URL")
	serviceRole := os.Getenv("SUPABASE_SERVICE_ROLE")
	client := CreateClient(url, serviceRole)

	user, err := client.Auth.GetUserFromEmail(context.Background(), "hunter@zapmoto.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(user.ID)
}