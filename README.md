# gowasm2dgame

golang port of [wxgame2](https://github.com/kasworld/wxgame2)

python 2.x 와 wxpython tcp connection 으로 만들었던 wxgame2 를 

golang, websocket, wasm 으로 포팅하는 프로젝트 입니다. 

python2가 지원이 종료된다니 겸사 겸사 이기도 합니다. (python3로 포팅할 계획은 없습니다.)

그동안 만들어 공개한 golang server program framework들 을 사용하는 예제 역할도 생각하고 있습니다.

[genprotocol](https://github.com/kasworld/genprotocol) 서버 클라이언트가 사용할 프로토콜 생성, 관리 

[argdefault](https://github.com/kasworld/argdefault) : config와 command line arguments 

[prettystring](https://github.com/kasworld/prettystring) : struct 의 string 화 / admin web , debug용 

[genenum](https://github.com/kasworld/genenum) : enum 의 생성, 관리 

[log](https://github.com/kasworld/log) : 전용 log package의 생성, 사용 

[signalhandle](https://github.com/kasworld/signalhandle) : signal을 관리해서 프로그램의 linux 서비스화, start,stop,forcestart,logreopen

정도가 중요하게 사용될 예정이며 그외에도 그동안 만들어둔 자잘한 package들을 적극 사용할 생각입니다. 

marshaler는 json, network은 websocket을 사용합니다. 

json 이 너무 느리고 패킷이 커져서 gob로 바꿨습니다. 
77kbyte 하던 것이 31kbyte로 주는 군요. 
(genprotocol을 사용하기에 수정해야할 코드는 단 네 줄입니다. -서버 2줄 클라이언트 2줄-  ^^)