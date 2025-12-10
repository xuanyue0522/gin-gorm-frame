package pool

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
)

/**
协程池
*/

type Pool struct {
	pool *ants.Pool
}

func NewPoolWithSize(size int) *Pool {
	pool, _ := ants.NewPool(size)
	return &Pool{
		pool: pool,
	}
}

func (p *Pool) RunGo(taskFun func()) {
	_ = p.pool.Submit(taskFun)
}

func (p *Pool) Wait() {
	p.pool.Release()
}

// 使用示例
func usePool() {
	pool := NewPoolWithSize(1)
	pool.RunGo(func() {
		fmt.Println("test 666")
		panic("test")
	})

	pool.RunGo(func() {
		fmt.Println("test 777")
		panic("test2")
	})

	pool.Wait()
}
