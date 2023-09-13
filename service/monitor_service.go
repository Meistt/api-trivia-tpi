package service

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type MonitorService interface {
	Observe(timeSelected int)
	TakeDate()
}

type monitorService struct {
	monitorService MonitorService
}

func NewMonitorService() MonitorService {
	return &monitorService{}
}

func (ms *monitorService) Observe(timeSelected int) {
	var (
		cpu      int
		mem      runtime.MemStats
		memUsage uint64
	)
	if timeSelected > 10 {
		for i := 0; i < timeSelected; i++ {
			cpu = runtime.NumGoroutine()
			runtime.ReadMemStats(&mem)
			memUsage = mem.Alloc / 1024
			fmt.Println("******************TIME******************")
			fmt.Println("-------->", i+1, "seg", " <--------")
			fmt.Printf("Uso de CPU: %d goroutines activas\n", cpu)
			fmt.Printf("Uso de Memoria: %d KB\n", memUsage)
			time.Sleep(time.Duration(i) * time.Second)
		}

	}
}

func (s *monitorService) TakeDate() {
	var (
		input    string
		intInput int
	)
	fmt.Println("****************** Bienvenido ******************")
	fmt.Println("Por favor ingrese un número entre 10 y 40:")
	time.Sleep(1 * time.Second)
	fmt.Println("**************************************************")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("No se pudo leer el valor ingresado: ", err)
		return
	}
	intInput, _ = strconv.Atoi(input)
	s.Observe(intInput)
}
