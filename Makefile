target=commit-per-day
$(target): main.go
	go build -o $(target) main.go

clean:
	rm $(target)