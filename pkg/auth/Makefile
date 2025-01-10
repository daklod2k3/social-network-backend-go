
dev:
	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command air -ErrorAction SilentlyContinue) { \
	air; \
	Write-Output 'Watching...'; \
} else { \
	Write-Output 'Installing air...'; \
	go install github.com/air-verse/air@latest; \
	air; \
	Write-Output 'Watching...'; \
}"

