package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// 创建上下文和取消函数
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 可选：设置无头模式和其他选项
	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true), // 使用无头模式
			chromedp.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36`),
		)...)
	defer cancelAlloc()
	ctx, cancel = chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// 定义变量来存储页面HTML
	var htmlContent string

	// 使用chromedp执行一系列任务
	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.ixigua.com/6726810241054278158`),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),             // 等待body元素可见
		chromedp.Sleep(5*time.Second),                              // 给予额外的时间让页面完全加载
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery), // 提取整个页面的HTML
	); err != nil {
		log.Fatal(err)
	}

	// 打印页面HTML内容
	fmt.Println("Page HTML:")
	fmt.Println(htmlContent)
}
