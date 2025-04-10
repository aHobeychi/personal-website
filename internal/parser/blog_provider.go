package parser

import (
	"aHobeychi/personal-website/internal/preprocessor"
)

// BlogProviderImpl implements the preprocessor.BlogProvider interface
type BlogProviderImpl struct{}

// GetAllBlogs returns all blogs
func (p *BlogProviderImpl) GetAllBlogs() ([]preprocessor.Blog, error) {
	blogs, err := ParseBlogs()
	if err != nil {
		return nil, err
	}

	result := make([]preprocessor.Blog, len(blogs))
	for i, blog := range blogs {
		result[i] = preprocessor.Blog{
			Id:    blog.Id,
			Title: blog.Title,
		}
	}

	return result, nil
}

// GetBlogContent returns the HTML content of a blog
func (p *BlogProviderImpl) GetBlogContent(blogId string) (string, error) {
	return GetBlogHTMLContent(blogId)
}

// GenerateAndSaveTableOfContents generates the table of contents for a blog post and saves it to a file
func (p *BlogProviderImpl) GenerateAndSaveTableOfContents(blogId string) error {
	blog, err := GetBlogByID(blogId)
	if err != nil {
		return err
	}

	content, err := GetBlogHTMLContent(blogId)
	if err != nil {
		return err
	}

	return preprocessor.GenerateAndSaveTableOfContents(preprocessor.Blog{
		Id:    blog.Id,
		Title: blog.Title,
	}, content)
}

// GetBlogProvider returns a new instance of BlogProviderImpl
func GetBlogProvider() preprocessor.BlogProvider {
	return &BlogProviderImpl{}
}
