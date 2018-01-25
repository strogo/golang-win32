package wevtapi

import "syscall"

var (
	wevtapi                              = syscall.NewLazyDLL("wevtapi.dll")
	evtArchiveExportedLog                = wevtapi.NewProc("EvtArchiveExportedLog")
	evtCancel                            = wevtapi.NewProc("EvtCancel")
	evtClearLog                          = wevtapi.NewProc("EvtClearLog")
	evtClose                             = wevtapi.NewProc("EvtClose")
	evtCreateBookmark                    = wevtapi.NewProc("EvtCreateBookmark")
	evtCreateRenderContext               = wevtapi.NewProc("EvtCreateRenderContext")
	evtExportLog                         = wevtapi.NewProc("EvtExportLog")
	evtFormatMessage                     = wevtapi.NewProc("EvtFormatMessage")
	evtGetChannelConfigProperty          = wevtapi.NewProc("EvtGetChannelConfigProperty")
	evtGetEventInfo                      = wevtapi.NewProc("EvtGetEventInfo")
	evtGetEventMetadataProperty          = wevtapi.NewProc("EvtGetEventMetadataProperty")
	evtGetExtendedStatus                 = wevtapi.NewProc("EvtGetExtendedStatus")
	evtGetLogInfo                        = wevtapi.NewProc("EvtGetLogInfo")
	evtGetObjectArrayProperty            = wevtapi.NewProc("EvtGetObjectArrayProperty")
	evtGetObjectArraySize                = wevtapi.NewProc("EvtGetObjectArraySize")
	evtGetPublisherMetadataProperty      = wevtapi.NewProc("EvtGetPublisherMetadataProperty")
	evtGetQueryInfo                      = wevtapi.NewProc("EvtGetQueryInfo")
	evtIntAssertConfig                   = wevtapi.NewProc("EvtIntAssertConfig")
	evtIntCreateBinXMLFromCustomXML      = wevtapi.NewProc("EvtIntCreateBinXMLFromCustomXML")
	evtIntCreateLocalLogfile             = wevtapi.NewProc("EvtIntCreateLocalLogfile")
	evtIntGetClassicLogDisplayName       = wevtapi.NewProc("EvtIntGetClassicLogDisplayName")
	evtIntRenderResourceEventTemplate    = wevtapi.NewProc("EvtIntRenderResourceEventTemplate")
	evtIntReportAuthzEventAndSourceAsync = wevtapi.NewProc("EvtIntReportAuthzEventAndSourceAsync")
	evtIntReportEventAndSourceAsync      = wevtapi.NewProc("EvtIntReportEventAndSourceAsync")
	evtIntRetractConfig                  = wevtapi.NewProc("EvtIntRetractConfig")
	evtIntSysprepCleanup                 = wevtapi.NewProc("EvtIntSysprepCleanup")
	evtIntWriteXmlEventToLocalLogfile    = wevtapi.NewProc("EvtIntWriteXmlEventToLocalLogfile")
	evtNextChannelPath                   = wevtapi.NewProc("EvtNextChannelPath")
	evtNextEventMetadata                 = wevtapi.NewProc("EvtNextEventMetadata")
	evtNextPublisherId                   = wevtapi.NewProc("EvtNextPublisherId")
	evtNext                              = wevtapi.NewProc("EvtNext")
	evtOpenChannelConfig                 = wevtapi.NewProc("EvtOpenChannelConfig")
	evtOpenChannelEnum                   = wevtapi.NewProc("EvtOpenChannelEnum")
	evtOpenEventMetadataEnum             = wevtapi.NewProc("EvtOpenEventMetadataEnum")
	evtOpenLog                           = wevtapi.NewProc("EvtOpenLog")
	evtOpenPublisherEnum                 = wevtapi.NewProc("EvtOpenPublisherEnum")
	evtOpenPublisherMetadata             = wevtapi.NewProc("EvtOpenPublisherMetadata")
	evtOpenSession                       = wevtapi.NewProc("EvtOpenSession")
	evtQuery                             = wevtapi.NewProc("EvtQuery")
	evtRender                            = wevtapi.NewProc("EvtRender")
	evtSaveChannelConfig                 = wevtapi.NewProc("EvtSaveChannelConfig")
	evtSeek                              = wevtapi.NewProc("EvtSeek")
	evtSetChannelConfigProperty          = wevtapi.NewProc("EvtSetChannelConfigProperty")
	evtSetObjectArrayProperty            = wevtapi.NewProc("EvtSetObjectArrayProperty")
	evtSubscribe                         = wevtapi.NewProc("EvtSubscribe")
	evtUpdateBookmark                    = wevtapi.NewProc("EvtUpdateBookmark")
)