#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <libgen.h>
#include <unistd.h>

int main(int argc, char* argv[]) {
	if (argc < 1)
		return 1;

	// Set user to 'root'
	setuid(0);

	// Clear env variables
	clearenv();

	// Check permissions
	struct stat info;
	stat(argv[1], &info);

	// Check if owner is root and permissions are set to 700 (= 0x1C0)
	printf("%x", info.st_mode);
	if (info.st_uid != 0 || (info.st_mode & 0x1FF) != 0x1C0) {
		printf("The deploy script has to be owned by root with permissions set to 700\n");
		return 2;
	}

	// Check directory permissions (to make sure the file can not be replaced)
	// TODO

	// Execute shell script
	system(argv[1]);

	return 0;
}
