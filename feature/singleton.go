package feature

import "sync"

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	if instance == nil {
		once.Do(func() {
			instance = &Singleton{}
		})
	}
	return instance
}

// 懒汉方式
// 缺点：非线程安全
func GetInstance_v1() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// 饿汉方式
// 缺点：如果初始化复杂，会影响加载时间；
func GetInstance_v2() *Singleton {
	// 在全局变量处就进行初始化，即instance = &Singleton{}
	return instance
}

var mu sync.Mutex

// 懒汉加锁方式
// 缺点：每次调用都需要加锁，影响性能
func GetInstance_v3() *Singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// 双重校验锁
func GetInstance_v4() *Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		// 此处再判断一次是否为空是为了防止两个线程同时通过了第一个if判断走到加锁流程，导致线程安全问题
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}
