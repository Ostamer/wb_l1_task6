package main

import (
	"context"
	"fmt"
	"time"
)

// Функция для заверешения горутины через закрытие канала
func closeChanel() {
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Горутина остановлена через закрытие канала")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(stop)
	time.Sleep(1 * time.Second)
}

// Функция для заверешения горутины через окончания цикла
func endCicle() {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Горутина завершена после 3 итераций")
	}()

	time.Sleep(5 * time.Second)
}

// Функция для заверешения горутины через окончания времени
func endTime() {
	go func() {
		timer := time.NewTimer(3 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				fmt.Println("Горутина остановлена через таймер")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
}

// Функция для заверешения горутины через контекст
func CloseContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена через context")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func main() {
	closeChanel()
	endTime()
	endCicle()
	CloseContext()

}
