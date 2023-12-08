
โดยย่อ Folder Structure ที่ใช้งานอยู่และกำลังจะ adjust จะไปในท่านี้ครับ
บาง project ก็จะตาม guide line ของแต่ล่ะที่ครับ

-cmd
  main.go // initial di, setup & connect DB, Start App.
-scr
  -core 
    - domain // keep entities which mapping or action with db
    - ports // keep interface
    - services // keep business logic that implemented working with handler and repository
  -repository // Keep repository struct that implemented and working with database
  -handler // handler func
  -error // Custom error in app
  -tests // Black box testing
-.env
- server
  - server.go // router initialize
go.mod
go.sum
docker-compose.yml
Makefile
