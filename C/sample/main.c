//
//  main.c
//  sqlitecloud-cli
//
//  Created by Marco Bambini on 08/02/21.
//

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sqcloud.h"
#include "linenoise.h"

#define HISTORY_FILENAME    ".sqlitecloud_history.txt"

void do_command (SQCloudConnection *conn, char *command) {
    SQCloudResult *res = SQCloudExec(conn, command);
    
    SQCloudResType type = SQCloudResultType(res);
    switch (type) {
        case RESULT_OK:
            printf("OK");
            break;
            
        case RESULT_ERROR:
            printf("ERROR: %s (%d)", SQCloudErrorMsg(conn), SQCloudErrorCode(conn));
            break;
            
        case RESULT_STRING:
            printf("%.*s", SQCloudResultLen(res), SQCloudResultBuffer(res));
            break;
            
        case RESULT_ROWSET:
            SQCloudRowSetDump(res);
            break;
    }
    
    printf("\n\n");
    SQCloudResultFree(res);
}

int main(int argc, const char * argv[]) {
    const char *hostname = "localhost";
    if (argc > 1) hostname = argv[1];
    
    SQCloudConnection *conn = SQCloudConnect(hostname, SQCLOUD_DEFAULT_PORT, NULL);
    if (SQCloudIsError(conn)) {
        printf("ERROR connecting to %s: %s (%d)\n", hostname, SQCloudErrorMsg(conn), SQCloudErrorCode(conn));
        return -1;
    } else {
        printf("Connection to %s OK...\n\n", hostname);
    }
    
    // load history file
    linenoiseHistoryLoad(HISTORY_FILENAME);
    
    // REPL
    char *command = NULL;
    while((command = linenoise(">> ")) != NULL) {
        if (command[0] != '\0') {
            linenoiseHistoryAdd(command);
            linenoiseHistorySave(HISTORY_FILENAME);
        }
        if (strncmp(command, "QUIT", 4) == 0) break;
        do_command(conn, command);
    }
    if (command) free(command);
    
    SQCloudFree(conn);
    return 0;
}
