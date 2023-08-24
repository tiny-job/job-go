package job

import "context"

type API interface {
	GetJobParams(ctx context.Context, pid int) (H, error)
	Next(ctx context.Context, pid int, resp H, err error) error
}

type Client struct {
	opts clientOptions
	apis API
}

// NewClient 创建一个客户端
func NewClient(opt ...ClientOption) *Client {
	opts := defaultClientOptions
	for _, o := range opt {
		o(&opts)
	}

	apis := NewApi(opts)

	return &Client{
		opts: opts,
		apis: apis,
	}
}

// GetJobParams 获取该任务当前参数
func (cli *Client) GetJobParams(ctx context.Context, pid int) (H, error) {
	return cli.apis.GetJobParams(ctx, pid)
}

// Next 下一步处理
func (cli *Client) Next(ctx context.Context, pid int, resp H, err error) error {
	return cli.apis.Next(ctx, pid, resp, err)
}
