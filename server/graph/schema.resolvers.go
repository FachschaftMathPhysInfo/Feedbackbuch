package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/FachschaftMathPhysInfo/Feedbackbuch/server/graph/generated"
	"github.com/FachschaftMathPhysInfo/Feedbackbuch/server/graph/model"
)

func (r *mutationResolver) CreateComment(ctx context.Context, content string, references *int) (*model.Comment, error) {
	comment := model.Comment{Content: content, Timestamp: time.Now(), References: references}
	if comment.Content == "" {
		return nil, fmt.Errorf("comment cannot be empty")
	}
	result := r.DB.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}

	go func() {
		r.mu.Lock()
		toBeDeleted := make([]int, 0)
		for i, c := range r.ToBeNotified {
			if c.Timestamp.Before(time.Now().Add(-20 * time.Minute)) {
				toBeDeleted = append(toBeDeleted, i)
			}
			c.C <- &comment
		}
		for i := range toBeDeleted {
			index := len(toBeDeleted) - 1 - i
			r.ToBeNotified = append(r.ToBeNotified[:index], r.ToBeNotified[index+1:]...)
		}
		r.mu.Unlock()
	}()

	return &comment, nil
}

func (r *mutationResolver) UpvoteComment(ctx context.Context, comment int) (*model.Comment, error) {
	upvotedComment := model.Comment{}
	r.DB.Where(&model.Comment{ID: comment}).First(&upvotedComment)
	if upvotedComment.ID == comment {
		upvotedComment.Upvotes++
		r.DB.Save(&upvotedComment)
		return &upvotedComment, nil
	}
	return nil, fmt.Errorf("The ID of the upvoted comment is invalid")
}

func (r *mutationResolver) DeleteComment(ctx context.Context, comment int) (*model.Comment, error) {
	deletedComment := model.Comment{}
	r.DB.Where(&model.Comment{ID: comment}).First(&deletedComment)

	if deletedComment.ID == comment {
		//Updates all references, where deleted comment is referenced
		deleted := -1
		r.DB.Model(&model.Comment{}).Where(&model.Comment{References: &comment}).Updates(model.Comment{References: &deleted})

		//Deletes original comment
		r.DB.Delete(&deletedComment)

		go func() {
			r.mu.Lock()
			toBeDeleted := make([]int, 0)
			for i, c := range r.ToBeNotified {
				if c.Timestamp.Before(time.Now().Add(-20 * time.Minute)) {
					toBeDeleted = append(toBeDeleted, i)
				}
				//BUG(tomtomrixrix): should not be sent to everyone.
				//c.C <- &deletedComment
			}
			for i := range toBeDeleted {
				index := len(toBeDeleted) - 1 + i
				r.ToBeNotified = append(r.ToBeNotified[:index], r.ToBeNotified[index+1:]...)
			}
			r.mu.Unlock()
		}()

		return &deletedComment, nil
	}

	return nil, fmt.Errorf("The ID of the deleted comment is invalid")
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	var results []*model.Comment
	r.DB.Find(&results)
	return results, nil
}

func (r *subscriptionResolver) CommentChanged(ctx context.Context) (<-chan *model.Comment, error) {
	event := Notification{C: make(chan *model.Comment), Timestamp: time.Now()}
	r.ToBeNotified = append(r.ToBeNotified, event)

	return event.C, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
