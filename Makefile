CC = go build

NAME = ArtilleryComputer.exe

SRCS = main.go

$(NAME):
	$(CC) $(SRCS)
	mv main.exe $(NAME)

all: $(NAME)

clean:
	rm -f $(NAME)

re: clean all