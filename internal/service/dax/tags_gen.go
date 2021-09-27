// Code generated by internal/generate/tags/main.go; DO NOT EDIT.

package dax

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dax"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// ListTags lists dax service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(conn *dax.DAX, identifier string) (tftags.KeyValueTags, error) {
	input := &dax.ListTagsInput{
		ResourceName: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return tftags.New(nil), err
	}

	return KeyValueTags(output.Tags), nil
}

// []*SERVICE.Tag handling

// Tags returns dax service tags.
func Tags(tags tftags.KeyValueTags) []*dax.Tag {
	result := make([]*dax.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &dax.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from dax service tags.
func KeyValueTags(tags []*dax.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(m)
}

// UpdateTags updates dax service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(conn *dax.DAX, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(oldTagsMap)
	newTags := tftags.New(newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &dax.UntagResourceInput{
			ResourceName: aws.String(identifier),
			TagKeys:      aws.StringSlice(removedTags.IgnoreAws().Keys()),
		}

		_, err := conn.UntagResource(input)

		if err != nil {
			return fmt.Errorf("error untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &dax.TagResourceInput{
			ResourceName: aws.String(identifier),
			Tags:         Tags(updatedTags.IgnoreAws()),
		}

		_, err := conn.TagResource(input)

		if err != nil {
			return fmt.Errorf("error tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
