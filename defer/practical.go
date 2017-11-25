func main() {
	file, err := os.Open("~/.bashrc")
	if err != nil {
		log.Fatalf("can't open: %v", err)
	}
	defer file.Close() // HL
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// do stuff
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("can't read: %v", err)
	}
}