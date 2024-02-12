package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"grpc/internal/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("could not connect to gRPC server: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := proto.NewBlogServiceClient(conn)

	fmt.Println("Choose operation: (1 for post, 2 for read, 3 for update, 4 for delete, -1 to stop)")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("invalid input: %v", err)
		}else{
			if choice == -1 {
				fmt.Println("Stopping client...")
				break
			}
	
			switch choice {
			case 1:
				fmt.Println("Enter post details:")
				fmt.Print("Title: ")
				scanner.Scan()
				title := scanner.Text()
				fmt.Print("Content: ")
				scanner.Scan()
				content := scanner.Text()
				fmt.Print("Author: ")
				scanner.Scan()
				author := scanner.Text()
				fmt.Print("Tags (comma-separated): ")
				scanner.Scan()
				tags := strings.Split(scanner.Text(), ",")
	
				createPostResponse, err := client.CreatePost(context.Background(), &proto.CreatePostRequest{
					Title:   title,
					Content: content,
					Author:  author,
					Tags:    tags,
				})
				if err != nil {
					fmt.Printf("error while calling CreatePost RPC: %v", err)
				}
				fmt.Printf("Response from CreatePost RPC: %v", createPostResponse)
	
			case 2:
				fmt.Println("Enter post ID to read:")
				scanner.Scan()
				postID := scanner.Text()
	
				readPostResponse, err := client.ReadPost(context.Background(), &proto.ReadPostRequest{
					PostId: postID,
				})
				if err != nil {
					fmt.Printf("error while calling ReadPost RPC: %v", err)
				}
				fmt.Printf("Response from ReadPost RPC: %v", readPostResponse)
	
			case 3:
				fmt.Println("Enter post ID to update:")
				scanner.Scan()
				postID := scanner.Text()
	
				fmt.Println("Enter updated post details:")
				fmt.Print("Title: ")
				scanner.Scan()
				title := scanner.Text()
				fmt.Print("Content: ")
				scanner.Scan()
				content := scanner.Text()
				fmt.Print("Author: ")
				scanner.Scan()
				author := scanner.Text()
				fmt.Print("Tags (comma-separated): ")
				scanner.Scan()
				tags := strings.Split(scanner.Text(), ",")
	
				updatePostResponse, err := client.UpdatePost(context.Background(), &proto.UpdatePostRequest{
					PostId:  postID,
					Title:   title,
					Content: content,
					Author:  author,
					Tags:    tags,
				})
				if err != nil {
					fmt.Printf("error while calling UpdatePost RPC: %v", err)
				}
				fmt.Printf("Response from UpdatePost RPC: %v", updatePostResponse)
	
			case 4:
				fmt.Println("Enter post ID to delete:")
				scanner.Scan()
				postID := scanner.Text()
	
				deletePostResponse, err := client.DeletePost(context.Background(), &proto.DeletePostRequest{
					PostId: postID,
				})
				if err != nil {
					fmt.Printf("error while calling DeletePost RPC: %v", err)
				}
				fmt.Printf("Response from DeletePost RPC: %v", deletePostResponse)

			case 5:
				fmt.Println("Reading all Post:")
				posts,err := client.ReadAllPost(context.Background(),&proto.ReadAllPostRequest{})
				if err != nil{
					fmt.Printf("error while reading all posts: %v",err)
				}
				fmt.Printf("Response for Read All RPC: %v",posts)
	
			default:
				fmt.Println("Invalid choice. Choose 1 for post, 2 for read, 3 for update, 4 for delete, or -1 to stop.")
			}
	
		}

		fmt.Println("Choose operation: (1 for post, 2 for read, 3 for update, 4 for delete, -1 to stop):")
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading standard input: %v", err)
		os.Exit(1)
	}
}
