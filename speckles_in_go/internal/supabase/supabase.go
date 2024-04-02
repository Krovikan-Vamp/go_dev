package supabase

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func InitSupabase() (*supabase.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	client, err := supabase.NewClient(os.Getenv("PUBLIC_SUPABASE_URL"), os.Getenv("PUBLIC_SUPABSE_CLIENT_API_KEY"), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
