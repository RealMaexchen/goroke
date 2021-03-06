package network

import (
	"net"
	"golang.org/x/net/websocket"
	"log"
	"fmt"
)



type NetworkController struct{
	IP net.IP
	Port uint16
	Connected bool

	Ws *websocket.Conn
}

func NewNetworkController(KaraFunIP net.IP, KaraFunPort uint16) *NetworkController{
	return &NetworkController{
		IP: KaraFunIP,
		Port:KaraFunPort,
		Connected:false,
		Ws:nil,
	}
}


func (self *NetworkController) GeneralVolumeSliderChanged(pos byte){
	self.sendString(fmt.Sprintf(GeneralVolumeCommand,pos))
}
func (self *NetworkController) VoiceVolumeSliderPosChanged(pos byte){
	self.sendString(fmt.Sprintf(VoiceVolumeCommand,pos))
}
func (self *NetworkController) MaleVolumeSliderPosChanged(pos byte){
	self.sendString(fmt.Sprintf(MaleVolumeCommand,pos))
}
func (self *NetworkController) FemaleVolumeSliderPosChanged(pos byte){
	self.sendString(fmt.Sprintf(FemaleVolumeCommand,pos))
}
func (self *NetworkController) PitchChanged(pitch int8){
	self.sendString(fmt.Sprintf(PitchCommand,pitch))
}
func (self *NetworkController) TempoChanged(tempo int8){
	self.sendString(fmt.Sprintf(TempoCommand, tempo))
}

func (self *NetworkController) PlayPressed(){
	self.sendString(PlayCommand)
}
func (self *NetworkController) PausePressed(){
	self.sendString(PauseCommand)
}
func (self *NetworkController) NextPressed(){
	self.sendString(NextCommand)
}
func (self *NetworkController) PrevPressed(){
	self.sendString(PrevCommand)
}



func (self *NetworkController) sendString(s string){
	if self.Connected && self.Ws != nil{
		if err:= websocket.Message.Send(self.Ws,s);err!=nil{
			log.Printf("cannot send: %v", err)
		}
	}
}

func (self *NetworkController) Connect() bool{
	if ip:=self.IP.To4();ip!= nil{
		dst:= fmt.Sprintf(url,self.IP.String(),self.Port)
		ws, err := websocket.Dial(dst, "", origin)
		if err != nil{
			self.Connected = false
			log.Printf("Cannot connect to the server: %v",dst)
			return false
		}

		self.Ws = ws
		log.Printf("Connected to the server: %v",dst)
		self.Connected = true
	}
	return true
}



