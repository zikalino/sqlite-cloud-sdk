package sqlitecloud

// #cgo CFLAGS: -Wno-multichar -I../../../C
// #cgo LDFLAGS: -L. -lsqcloud -ldl
// #include <stdlib.h>
// #include "sqcloud.h"
import "C"
import "unsafe"
//import "errors"
// import "fmt"
// import "reflect"

// SQCloudConnection *SQCloudConnect (const char *hostname, int port, SQCloudConfig *config);
func CConnect( Host string, Port int, Username string, Password string, Database string, Timeout int, Family int ) *C.struct_SQCloudConnection {
  conf := C.struct_SQCloudConfigStruct{}
  conf.username = C.CString( Username )
  conf.password = C.CString( Password )
  conf.database = C.CString( Database )
  conf.timeout  = C.int( Timeout )
  conf.family   = C.int( Family )

  cHost := C.CString( Host )

  cConnection := C.SQCloudConnect( cHost, C.int( Port ), &conf )
  
  C.free( unsafe.Pointer( cHost ) )
  C.free( unsafe.Pointer( conf.database ) )
  C.free( unsafe.Pointer( conf.password ) )
  C.free( unsafe.Pointer( conf.username ) )

  return cConnection
}

// SQCloudConnection *SQCloudConnectWithString (const char *s);
func CConnectWithString( ConnectionString string ) *SQCloud {
  cConString := C.CString( ConnectionString )
  connection := SQCloud{ connection: C.SQCloudConnectWithString( cConString ) }
  C.free( unsafe.Pointer( cConString ) )

  if connection.connection == nil {
    return nil
  }

  return &connection
}

// void SQCloudDisconnect (SQCloudConnection *connection);
func (this *SQCloud ) CDisconnect() {
  if this.connection != nil {
    C.SQCloudDisconnect( this.connection )
    this.connection = nil
  }
}
// char *SQCloudUUID (SQCloudConnection *connection);
func (this *SQCloud ) CGetCloudUUID() string {
   return C.GoString( C.SQCloudUUID( this.connection ) )
}

//bool SQCloudIsError (SQCloudConnection *connection);
func (this *SQCloud ) CIsError() bool {
  return bool( C.SQCloudIsError( this.connection ) )
}
//int SQCloudErrorCode (SQCloudConnection *connection);
func (this *SQCloud ) CGetErrorCode() int {
  return int( C.SQCloudErrorCode( this.connection ) )
}
//const char *SQCloudErrorMsg (SQCloudConnection *connection);
func (this *SQCloud ) CGetErrorMessage() string {
  return C.GoString( C.SQCloudErrorMsg( this.connection ) )
}

// SQCloudResult *SQCloudExec (SQCloudConnection *connection, const char *command);
func (this *SQCloud ) CExec( Command string ) *SQCloudResult {
  cCommand := C.CString( Command )
  defer C.free( unsafe.Pointer( cCommand ) )

  // println( "exec ("+Command+").." )

  result := SQCloudResult{ result: C.SQCloudExec( this.connection, cCommand ) }
  if result.result == nil {
    return nil
  }
  return &result
}
// SQCloudResult *SQCloudSetPubSubOnly (SQCloudConnection *connection);
func (this *SQCloud ) CSetPubSubOnly() *SQCloudResult {
  result := SQCloudResult{ result: C.SQCloudSetPubSubOnly( this.connection ) }
  
  if result.result == nil {
    return nil
  }

  return &result
}
// SQCloudResType SQCloudResultType (SQCloudResult *result);
func (this *SQCloudResult ) CGetResultType() uint {
  return uint( C.SQCloudResultType( this.result ) )
}
// uint32_t SQCloudResultLen (SQCloudResult *result);
func (this *SQCloudResult ) CGetResultLen() uint {
  return uint( C.SQCloudResultLen( this.result ) )
}
// char *SQCloudResultBuffer (SQCloudResult *result);
func (this *SQCloudResult ) CGetResultBuffer() string {
  return C.GoString( C.SQCloudResultBuffer( this.result ) )
}
// void SQCloudResultFree (SQCloudResult *result);
func (this *SQCloudResult ) CFree() {
  C.SQCloudResultFree( this.result )
}
// bool SQCloudResultIsOK (SQCloudResult *result);
func (this *SQCloudResult ) CIsOK() bool {
  return bool( C.SQCloudResultIsOK( this.result ) )
}
// SQCloudValueType SQCloudRowsetValueType (SQCloudResult *result, uint32_t row, uint32_t col);
func (this *SQCloudResult ) CGetValueType( Row uint, Column uint ) int {
  return int( C.SQCloudRowsetValueType( this.result, C.uint( Row ), C.uint( Column ) ) )
}
// uint32_t SQCloudResultMaxColumnLenght (SQCloudResult *result, uint32_t col) ;
func (this *SQCloudResult ) CGetMaxColumnLenght( Column uint ) uint {
  return uint( C.SQCloudRowsetRowsMaxColumnLength( this.result, C.uint( Column ) ) )
}
// char *SQCloudRowsetColumnName (SQCloudResult *result, uint32_t col, uint32_t *len);
func (this *SQCloudResult ) CGetColumnName( Column uint ) string {
  var len C.uint32_t = 0
  return C.GoStringN( C.SQCloudRowsetColumnName( this.result, C.uint( Column ), &len ), C.int( len ) )
}
// uint32_t SQCloudRowsetRows (SQCloudResult *result);
func (this *SQCloudResult ) CGetRows() uint {
  return uint( C.SQCloudRowsetRows( this.result ) )
}
// uint32_t SQCloudRowsetCols (SQCloudResult *result);
func (this *SQCloudResult ) CGetColumns() uint {
  return uint( C.SQCloudRowsetCols( this.result ) )
}
// uint32_t SQCloudRowsetMaxLen (SQCloudResult *result);
func (this *SQCloudResult ) CGetMaxLen() uint32 {
  return uint32( C.SQCloudRowsetMaxLen( this.result ) )
}
// char *SQCloudRowsetValue (SQCloudResult *result, uint32_t row, uint32_t col, uint32_t *len);
func (this *SQCloudResult ) CGetStringValue( Row uint, Column uint ) string {
  var len C.uint32_t = 0
  return C.GoStringN( C.SQCloudRowsetValue( this.result, C.uint32_t( Row ), C.uint32_t( Column ), &len ), C.int( len ) ) // Problem: NULL Pointer in return
}
// int32_t SQCloudRowsetInt32Value (SQCloudResult *result, uint32_t row, uint32_t col);
func (this *SQCloudResult ) CGetInt32Value( Row uint, Column uint ) int32 {
  return int32( C.SQCloudRowsetInt32Value( this.result, C.uint( Row ), C.uint( Column ) ) )
}
// int64_t SQCloudRowsetInt64Value (SQCloudResult *result, uint32_t row, uint32_t col);
func (this *SQCloudResult ) CGetInt64Value( Row uint, Column uint ) int64 {
  return int64( C.SQCloudRowsetInt64Value( this.result, C.uint( Row ), C.uint( Column ) ) )
}
// float SQCloudRowsetFloatValue (SQCloudResult *result, uint32_t row, uint32_t col);
func (this *SQCloudResult ) CGetFloat32Value( Row uint, Column uint ) float32 {
  return float32( C.SQCloudRowsetFloatValue( this.result, C.uint( Row ), C.uint( Column ) ) )
}
// double SQCloudRowsetDoubleValue (SQCloudResult *result, uint32_t row, uint32_t col);
func (this *SQCloudResult ) CGetFloat64Value( Row uint, Column uint ) float64 {
  return float64( C.SQCloudRowsetDoubleValue( this.result, C.uint( Row ), C.uint( Column ) ) )
}
// void SQCloudRowsetDump (SQCloudResult *result, uint32_t maxline);
func (this *SQCloudResult ) CDump( MaxLine uint ) {
   C.SQCloudRowsetDump( this.result, C.uint( MaxLine ) )
}

// Reserverd (internal) functions - will never be exported

// bool SQCloudForwardExec(SQCloudConnection *connection, const char *command, bool (*forward_cb) (char *buffer, size_t blen, void *xdata), void *xdata) {
// SQCloudResult *SQCloudSetUUID (SQCloudConnection *connection, const char *UUID) 

// Will be implemented in GO

// void SQCloudSetPubSubCallback (SQCloudConnection *connection, SQCloudPubSubCB callback, void *data);