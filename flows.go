package main

//
//import (
//	"github.com/TIBCOSoftware/flogo-lib/core/action"
//	"github.com/TIBCOSoftware/flogo-lib/engine"
//	"github.com/TIBCOSoftware/flogo-lib/flow/flowinst"
//	"github.com/TIBCOSoftware/flogo-lib/flow/service"
//	"github.com/TIBCOSoftware/flogo-lib/flow/service/flowprovider"
//	"github.com/TIBCOSoftware/flogo-lib/flow/service/staterecorder"
//	"github.com/TIBCOSoftware/flogo-lib/flow/service/tester"
//	"github.com/TIBCOSoftware/flogo-lib/flow/support"
//)
//
//var embeddedJSONFlows map[string]string
//
//func init() {
//
//	embeddedJSONFlows = make(map[string]string)
//
//	embeddedJSONFlows["embedded://myflow"] = "H4sIAAAJbogC/7yTMW+DMBCF9/yKk+d0SEe2dqjUtcpWZTD4kp4wvtQ+qKIo/70CnGCEVUVVVTbee777ZB7nFQCA0iKeylYwqALed+tRbdigVQUoobLih0DN0aKKptMN9t7e8tdV88yy1aFWBYyDx+GVUEdy2p6Ow4mYHjwyqoBNIlhydQox29XPT0+LDmP4JvXPefaWRSgtV3X1ockl46b07DYWfn7HArbBEPQBMwsm/ogTxJM7/JTstG2H6BtWSB0a8BgEPH62GERlD17Wv0Xvv+mr2/M97CWzRe3ughff4l+zamO2/DKV8D9gF+ou06Kh2o8Z40r+fCshPMV+5toY2TerPEL6n8yi8Uon8fINAAD//wEAAP//11Y5w+0DAAA="
//
//}
//
//// EnableFlowServices enables flow services and action for engine
//func EnableFlowServices(engine *engine.Engine, engineConfig *engine.Config) {
//
//	embeddedFlowMgr := support.NewEmbeddedFlowManager(true, embeddedJSONFlows)
//
//	fpConfig := engineConfig.Services[service.ServiceFlowProvider]
//	flowProvider := flowprovider.NewRemoteFlowProvider(fpConfig, embeddedFlowMgr)
//	engine.RegisterService(flowProvider)
//
//	srConfig := engineConfig.Services[service.ServiceStateRecorder]
//	stateRecorder := staterecorder.NewRemoteStateRecorder(srConfig)
//	engine.RegisterService(stateRecorder)
//
//	etConfig := engineConfig.Services[service.ServiceEngineTester]
//	engineTester := tester.NewRestEngineTester(etConfig)
//	engine.RegisterService(engineTester)
//
//	options := &flowinst.ActionOptions{Record: stateRecorder.Enabled()}
//
//	flowAction := flowinst.NewFlowAction(flowProvider, stateRecorder, options)
//	action.Register(flowinst.ActionType, flowAction)
//}
