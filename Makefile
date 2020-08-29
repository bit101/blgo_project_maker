default: clean
	@mkdir frames
	@go run main.go


clean:
	@rm -rf frames
	@rm -f out.png
	@rm -f out.gif
