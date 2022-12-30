package pulumiplaygroundlib

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SimplePublicS3FileShare struct {
	pulumi.ResourceState
	BucketName pulumi.IDOutput
}

func NewMyComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*SimplePublicS3FileShare, error) {
	share := &SimplePublicS3FileShare{}

	err := ctx.RegisterComponentResource("nerja:pulumi:SimplePublicS3FileShare", name, share, opts...)
	if err != nil {
		return nil, err
	}

	bucket, err := s3.NewBucket(ctx, fmt.Sprintf("%s-bucket", name), &s3.BucketArgs{
		Website: s3.BucketWebsiteArgs{
			IndexDocument: pulumi.String("index.html"),
		},
		Acl: pulumi.String("public-read"),
	}, pulumi.Parent(share))
	if err != nil {
		return nil, err
	}

	_, err = s3.NewBucketObject(ctx, "obj", &s3.BucketObjectArgs{
		Bucket:  bucket.ID(),
		Content: pulumi.String("hello"),
		Key:     pulumi.String("index.html"),
		Acl:     pulumi.String("public-read"),
	}, pulumi.Parent(bucket))
	if err != nil {
		return nil, err
	}

	share.BucketName = bucket.ID()

	ctx.RegisterResourceOutputs(share, pulumi.Map{
		"bucketName": bucket.ID(),
	})

	return share, nil
}
