package supabase

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func InitSupabase() (*supabase.Client, error) {
	err := godotenv.Load()
	fmt.Println(os.Getenv("PUBLIC_SUPABASE_URL"), "\n\n", os.Getenv("PUBLIC_SUPABSE_CLIENT_API_KEY"))
	if err != nil {
		return nil, err
	}

	client, err := supabase.NewClient(os.Getenv("PUBLIC_SUPABASE_URL"), os.Getenv("PUBLIC_SUPABSE_CLIENT_API_KEY"), nil)

	data, count, err := client.From("profiles").Select("*", "exact", false).Execute()

	fmt.Println(string(data), err, count)

	if err != nil {
		return nil, err
	}

	return client, nil
}
