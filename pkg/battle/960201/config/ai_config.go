package config
var AiConfigStr = `[{
  "IntoRoom1Person": 100,
  "IntoRoom2Person": 80,
  "IntoRoom3Person": 50,
  "LeaveRoom5Person": 100,
  "LeaveRoom3Person": 0,

  "FirstRoundSee": 0,
  "FirstRoundFollow": 80,
  "FirstRoundRaise": 20,
  "FirstRoundCompare": 0,
  "FirstRoundAllIn": 0,

  "SecondRoundNotSee1": 5,
  "SecondRoundNotSee2": 10,
  "SecondRoundNotSeeOver2": 30,
  "SecondRoundNotSeeGiveUp": 0,
  "SecondRoundNotSeeFollow": 80,
  "SecondRoundNotSeeRaise": 20,
  "SecondRoundNotSeeCompare": 0,
  "SecondRoundNotSeeAllin": 0,
  "SecondRoundSawOverKGiveUp": 0,
  "SecondRoundSawBelowKGiveUp": 80,
  "SecondRoundSawFollow": 70,
  "SecondRoundSawRaise": 30,
  "SecondRoundSawCompare": 0,
  "SecondRoundSawAllin": 0,

  "ThirdRoundNotSee1": 20,
  "ThirdRoundNotSee2": 30,
  "ThirdRoundNotSeeOver2": 40,
  "ThirdRoundNotSeeGiveUp": 0,
  "ThirdRoundNotSeeFollow": 75,
  "ThirdRoundNotSeeRaise": 20,
  "ThirdRoundNotSeeCompare": 5,
  "ThirdRoundNotSeeAllin": 0,

  "ThirdRoundSawBelowKGiveUp": 90,
  "ThirdRoundDzGiveUp": 10,
  "ThirdRoundDzFollow": 65,
  "ThirdRoundDzRaise": 20,
  "ThirdRoundDzCompare": 10,
  "ThirdRoundDzAllin": 5,

  "ThirdRoundSzGiveUp": 0,
  "ThirdRoundSzAllin": 5,
  "ThirdRoundSzCompare": 10,
  "ThirdRoundSzRaise": 40,
  "ThirdRoundSzFollow": 45,

  "ThirdRoundJhGiveUp": 0,
  "ThirdRoundJhFollow": 30,
  "ThirdRoundJhRaise": 60,
  "ThirdRoundJhCompare": 5,
  "ThirdRoundJhAllin": 5,

  "ThirdRoundOverJhGiveUp": 0,
  "ThirdRoundOverJhFollow": 0,
  "ThirdRoundOverJhRaise": 90,
  "ThirdRoundOverJhCompare": 0,
  "ThirdRoundOverJhAllin": 10,

  "FourOverNotSee1":70,
  "FourOverNotSee2": 70,
  "FourOverNotSeeOver2":70,
  "FourOverNotSeeGiveUp":0,
  "FourOverNotSeeFollow":75,
  "FourOverNotSeeRaise":20,
  "FourOverNotSeeCompare":5,
  "FourOverNotSeeAllin":0,

  "FourOverBelowKGiveUp":100,
  "FourOverDzGiveUp":10,
  "FourOverDzFollow":70,
  "FourOverDzRaise":20,
  "FourOverDzCompare":10,
  "FourOverDzAllin":0,

  "FourOverSzGiveUp":0,
  "FourOverSzAllin":5,
  "FourOverSzCompare":10,
  "FourOverSzRaise":40,
  "FourOverSzFollow":45,

  "FourOverJhGiveUp":0,
  "FourOverJhFollow":30,
  "FourOverJhRaise":60,
  "FourOverJhCompare":5,
  "FourOverJhAllin":5,

  "FourOverOverJhGiveUp":0,
  "FourOverOverJhFollow":30,
  "FourOverOverJhRaise":50,
  "FourOverOverJhCompare":0,
  "FourOverOverJhAllin":20,


  "FirstRoundRaise1": 30,
  "FirstRoundRaise2": 50,
  "FirstRoundRaise3": 20,
  "SecondRoundNotSeeRaise1": 30,
  "SecondRoundNotSeeRaise2": 50,
  "SecondRoundNotSeeRaise3": 20,
  "SecondRoundSawRaise1": 30,
  "SecondRoundSawRaise2": 50,
  "SecondRoundSawRaise3": 20,
  "ThirdRoundNotSeeRaise1": 30,
  "ThirdRoundNotSeeRaise2": 50,
  "ThirdRoundNotSeeRaise3": 20,
  "ThirdRoundDzRaise1": 30,
  "ThirdRoundDzRaise2": 50,
  "ThirdRoundDzRaise3": 20,
  "ThirdRoundSzRaise1": 10,
  "ThirdRoundSzRaise2": 50,
  "ThirdRoundSzRaise3": 40,
  "ThirdRoundJhRaise1": 10,
  "ThirdRoundJhRaise2": 30,
  "ThirdRoundJhRaise3": 60,
  "ThirdRoundOverJhRaise1": 10,
  "ThirdRoundOverJhRaise2": 30,
  "ThirdRoundOverJhRaise3": 60,
  "FourRoundNotSeeRaise1": 30,
  "FourRoundNotSeeRaise2": 50,
  "FourRoundNotSeeRaise3": 20,
  "FourRoundDzRaise1": 30,
  "FourRoundDzRaise2": 50,
  "FourRoundDzRaise3": 20,
  "FourRoundSzRaise1": 10,
  "FourRoundSzRaise2": 50,
  "FourRoundSzRaise3": 40,
  "FourRoundJhRaise1": 10,
  "FourRoundJhRaise2": 30,
  "FourRoundJhRaise3": 60,
  "FourRoundOverJhRaise1": 10,
  "FourRoundOverJhRaise2": 30,
  "FourRoundOverJhRaise3": 60
},
  {
  "IntoRoom1Person": 100,
  "IntoRoom2Person": 80,
  "IntoRoom3Person": 50,
  "LeaveRoom5Person": 100,
  "LeaveRoom3Person": 0,

  "FirstRoundSee": 0,
  "FirstRoundFollow": 60,
  "FirstRoundRaise": 40,
  "FirstRoundCompare": 0,
  "FirstRoundAllIn": 0,

  "SecondRoundNotSee1": 60,
  "SecondRoundNotSee2": 70,
  "SecondRoundNotSeeOver2": 80,
  "SecondRoundNotSeeGiveUp": 0,
  "SecondRoundNotSeeFollow": 60,
  "SecondRoundNotSeeRaise": 40,
  "SecondRoundNotSeeCompare": 0,
  "SecondRoundNotSeeAllin": 0,
  "SecondRoundSawOverKGiveUp": 0,
  "SecondRoundSawBelowKGiveUp": 30,
  "SecondRoundSawFollow": 80,
  "SecondRoundSawRaise": 20,
  "SecondRoundSawCompare": 0,
  "SecondRoundSawAllin": 0,

  "ThirdRoundNotSee1": 70,
  "ThirdRoundNotSee2": 80,
  "ThirdRoundNotSeeOver2": 90,
  "ThirdRoundNotSeeGiveUp": 0,
  "ThirdRoundNotSeeFollow": 90,
  "ThirdRoundNotSeeRaise": 5,
  "ThirdRoundNotSeeCompare": 5,
  "ThirdRoundNotSeeAllin": 0,

  "ThirdRoundSawBelowKGiveUp": 30,
  "ThirdRoundDzGiveUp": 10,
  "ThirdRoundDzFollow": 75,
  "ThirdRoundDzRaise": 10,
  "ThirdRoundDzCompare": 10,
  "ThirdRoundDzAllin": 5,

  "ThirdRoundSzGiveUp": 0,
  "ThirdRoundSzAllin": 5,
  "ThirdRoundSzCompare": 10,
  "ThirdRoundSzRaise": 30,
  "ThirdRoundSzFollow": 55,

  "ThirdRoundJhGiveUp": 0,
  "ThirdRoundJhFollow": 40,
  "ThirdRoundJhRaise": 50,
  "ThirdRoundJhCompare": 5,
  "ThirdRoundJhAllin": 5,

  "ThirdRoundOverJhGiveUp": 0,
  "ThirdRoundOverJhFollow": 10,
  "ThirdRoundOverJhRaise": 80,
  "ThirdRoundOverJhCompare": 0,
  "ThirdRoundOverJhAllin": 10,

  "FourOverNotSee1":100,
  "FourOverNotSee2": 100,
  "FourOverNotSeeOver2":100,
  "FourOverNotSeeGiveUp":0,
  "FourOverNotSeeFollow":90,
  "FourOverNotSeeRaise":5,
  "FourOverNotSeeCompare":5,
  "FourOverNotSeeAllin":0,

  "FourOverBelowKGiveUp":100,
  "FourOverDzGiveUp":10,
  "FourOverDzFollow":80,
  "FourOverDzRaise":10,
  "FourOverDzCompare":10,
  "FourOverDzAllin":0,

  "FourOverSzGiveUp":0,
  "FourOverSzAllin":5,
  "FourOverSzCompare":5,
  "FourOverSzRaise":30,
  "FourOverSzFollow":60,

  "FourOverJhGiveUp":0,
  "FourOverJhFollow":40,
  "FourOverJhRaise":50,
  "FourOverJhCompare":5,
  "FourOverJhAllin":5,

  "FourOverOverJhGiveUp":0,
  "FourOverOverJhFollow":40,
  "FourOverOverJhRaise":40,
  "FourOverOverJhCompare":0,
  "FourOverOverJhAllin":20,


  "FirstRoundRaise1": 60,
  "FirstRoundRaise2": 30,
  "FirstRoundRaise3": 10,
  "SecondRoundNotSeeRaise1": 10,
  "SecondRoundNotSeeRaise2": 50,
  "SecondRoundNotSeeRaise3": 40,
  "SecondRoundSawRaise1": 10,
  "SecondRoundSawRaise2": 50,
  "SecondRoundSawRaise3": 40,
  "ThirdRoundNotSeeRaise1": 60,
  "ThirdRoundNotSeeRaise2": 30,
  "ThirdRoundNotSeeRaise3": 10,
  "ThirdRoundDzRaise1": 60,
  "ThirdRoundDzRaise2": 30,
  "ThirdRoundDzRaise3": 10,
  "ThirdRoundSzRaise1": 10,
  "ThirdRoundSzRaise2": 60,
  "ThirdRoundSzRaise3": 30,
  "ThirdRoundJhRaise1": 10,
  "ThirdRoundJhRaise2": 30,
  "ThirdRoundJhRaise3": 60,
  "ThirdRoundOverJhRaise1": 10,
  "ThirdRoundOverJhRaise2": 30,
  "ThirdRoundOverJhRaise3": 60,
  "FourRoundNotSeeRaise1": 60,
  "FourRoundNotSeeRaise2": 30,
  "FourRoundNotSeeRaise3": 10,
  "FourRoundDzRaise1": 60,
  "FourRoundDzRaise2": 30,
  "FourRoundDzRaise3": 10,
  "FourRoundSzRaise1": 10,
  "FourRoundSzRaise2": 60,
  "FourRoundSzRaise3": 30,
  "FourRoundJhRaise1": 10,
  "FourRoundJhRaise2": 30,
  "FourRoundJhRaise3": 60,
  "FourRoundOverJhRaise1": 10,
  "FourRoundOverJhRaise2": 30,
  "FourRoundOverJhRaise3": 60
},
  {
  "IntoRoom1Person": 100,
  "IntoRoom2Person": 80,
  "IntoRoom3Person": 50,
  "LeaveRoom5Person": 100,
  "LeaveRoom3Person": 0,

  "FirstRoundSee": 0,
  "FirstRoundFollow": 95,
  "FirstRoundRaise": 5,
  "FirstRoundCompare": 0,
  "FirstRoundAllIn": 0,

  "SecondRoundNotSee1": 60,
  "SecondRoundNotSee2": 80,
  "SecondRoundNotSeeOver2": 100,
  "SecondRoundNotSeeGiveUp": 0,
  "SecondRoundNotSeeFollow": 95,
  "SecondRoundNotSeeRaise": 5,
  "SecondRoundNotSeeCompare": 0,
  "SecondRoundNotSeeAllin": 0,
  "SecondRoundSawOverKGiveUp": 0,
  "SecondRoundSawBelowKGiveUp": 100,
  "SecondRoundSawFollow": 90,
  "SecondRoundSawRaise": 10,
  "SecondRoundSawCompare": 0,
  "SecondRoundSawAllin": 0,

  "ThirdRoundNotSee1": 80,
  "ThirdRoundNotSee2": 90,
  "ThirdRoundNotSeeOver2": 100,
  "ThirdRoundNotSeeGiveUp": 0,
  "ThirdRoundNotSeeFollow": 90,
  "ThirdRoundNotSeeRaise": 5,
  "ThirdRoundNotSeeCompare": 5,
  "ThirdRoundNotSeeAllin": 0,

  "ThirdRoundSawBelowKGiveUp": 100,
  "ThirdRoundDzGiveUp": 50,
  "ThirdRoundDzFollow": 75,
  "ThirdRoundDzRaise": 10,
  "ThirdRoundDzCompare": 10,
  "ThirdRoundDzAllin": 5,

  "ThirdRoundSzGiveUp":10,
  "ThirdRoundSzAllin": 5,
  "ThirdRoundSzCompare": 10,
  "ThirdRoundSzRaise": 30,
  "ThirdRoundSzFollow": 55,

  "ThirdRoundJhGiveUp": 5,
  "ThirdRoundJhFollow": 40,
  "ThirdRoundJhRaise": 50,
  "ThirdRoundJhCompare": 5,
  "ThirdRoundJhAllin": 5,

  "ThirdRoundOverJhGiveUp": 0,
  "ThirdRoundOverJhFollow": 10,
  "ThirdRoundOverJhRaise": 80,
  "ThirdRoundOverJhCompare": 0,
  "ThirdRoundOverJhAllin": 10,

  "FourOverNotSee1":100,
  "FourOverNotSee2": 100,
  "FourOverNotSeeOver2":100,
  "FourOverNotSeeGiveUp":0,
  "FourOverNotSeeFollow":90,
  "FourOverNotSeeRaise":5,
  "FourOverNotSeeCompare":5,
  "FourOverNotSeeAllin":0,

  "FourOverBelowKGiveUp":100,
  "FourOverDzGiveUp":100,
  "FourOverDzFollow":80,
  "FourOverDzRaise":10,
  "FourOverDzCompare":10,
  "FourOverDzAllin":0,

  "FourOverSzGiveUp":10,
  "FourOverSzAllin":5,
  "FourOverSzCompare":10,
  "FourOverSzRaise":30,
  "FourOverSzFollow":55,

  "FourOverJhGiveUp":5,
  "FourOverJhFollow":40,
  "FourOverJhRaise":50,
  "FourOverJhCompare":5,
  "FourOverJhAllin":5,

  "FourOverOverJhGiveUp":0,
  "FourOverOverJhFollow":40,
  "FourOverOverJhRaise":40,
  "FourOverOverJhCompare":0,
  "FourOverOverJhAllin":20,


  "FirstRoundRaise1": 80,
  "FirstRoundRaise2": 15,
  "FirstRoundRaise3": 5,
  "SecondRoundNotSeeRaise1": 80,
  "SecondRoundNotSeeRaise2": 15,
  "SecondRoundNotSeeRaise3": 5,
  "SecondRoundSawRaise1": 80,
  "SecondRoundSawRaise2": 15,
  "SecondRoundSawRaise3": 5,
  "ThirdRoundNotSeeRaise1": 80,
  "ThirdRoundNotSeeRaise2": 15,
  "ThirdRoundNotSeeRaise3": 5,
  "ThirdRoundDzRaise1": 80,
  "ThirdRoundDzRaise2": 15,
  "ThirdRoundDzRaise3": 5,
  "ThirdRoundSzRaise1": 10,
  "ThirdRoundSzRaise2": 60,
  "ThirdRoundSzRaise3": 30,
  "ThirdRoundJhRaise1": 10,
  "ThirdRoundJhRaise2": 30,
  "ThirdRoundJhRaise3": 60,
  "ThirdRoundOverJhRaise1": 10,
  "ThirdRoundOverJhRaise2": 30,
  "ThirdRoundOverJhRaise3": 60,
  "FourRoundNotSeeRaise1": 80,
  "FourRoundNotSeeRaise2": 15,
  "FourRoundNotSeeRaise3": 5,
  "FourRoundDzRaise1": 80,
  "FourRoundDzRaise2": 15,
  "FourRoundDzRaise3": 5,
  "FourRoundSzRaise1": 10,
  "FourRoundSzRaise2": 60,
  "FourRoundSzRaise3": 30,
  "FourRoundJhRaise1": 10,
  "FourRoundJhRaise2": 30,
  "FourRoundJhRaise3": 60,
  "FourRoundOverJhRaise1": 10,
  "FourRoundOverJhRaise2": 30,
  "FourRoundOverJhRaise3": 60
  },
  {
  "IntoRoom1Person": 100,
  "IntoRoom2Person": 80,
  "IntoRoom3Person": 50,
  "LeaveRoom5Person": 100,
  "LeaveRoom3Person": 0,

  "FirstRoundSee": 0,
  "FirstRoundFollow": 95,
  "FirstRoundRaise": 5,
  "FirstRoundCompare": 0,
  "FirstRoundAllIn": 0,

  "SecondRoundNotSee1": 5,
  "SecondRoundNotSee2": 20,
  "SecondRoundNotSeeOver2": 50,
  "SecondRoundNotSeeGiveUp": 0,
  "SecondRoundNotSeeFollow": 95,
  "SecondRoundNotSeeRaise": 5,
  "SecondRoundNotSeeCompare": 0,
  "SecondRoundNotSeeAllin": 0,
  "SecondRoundSawOverKGiveUp": 0,
  "SecondRoundSawBelowKGiveUp": 80,
  "SecondRoundSawFollow": 90,
  "SecondRoundSawRaise": 10,
  "SecondRoundSawCompare": 0,
  "SecondRoundSawAllin": 0,

  "ThirdRoundNotSee1": 20,
  "ThirdRoundNotSee2": 40,
  "ThirdRoundNotSeeOver2": 80,
  "ThirdRoundNotSeeGiveUp": 0,
  "ThirdRoundNotSeeFollow": 90,
  "ThirdRoundNotSeeRaise": 5,
  "ThirdRoundNotSeeCompare": 5,
  "ThirdRoundNotSeeAllin": 0,

  "ThirdRoundSawBelowKGiveUp": 90,
  "ThirdRoundDzGiveUp": 10,
  "ThirdRoundDzFollow": 75,
  "ThirdRoundDzRaise": 10,
  "ThirdRoundDzCompare": 10,
  "ThirdRoundDzAllin": 5,

  "ThirdRoundSzGiveUp": 0,
  "ThirdRoundSzAllin": 5,
  "ThirdRoundSzCompare": 10,
  "ThirdRoundSzRaise": 30,
  "ThirdRoundSzFollow": 55,

  "ThirdRoundJhGiveUp": 0,
  "ThirdRoundJhFollow": 40,
  "ThirdRoundJhRaise": 50,
  "ThirdRoundJhCompare": 5,
  "ThirdRoundJhAllin": 5,

  "ThirdRoundOverJhGiveUp": 0,
  "ThirdRoundOverJhFollow": 10,
  "ThirdRoundOverJhRaise": 80,
  "ThirdRoundOverJhCompare": 0,
  "ThirdRoundOverJhAllin": 10,

  "FourOverNotSee1":90,
  "FourOverNotSee2": 90,
  "FourOverNotSeeOver2":90,
  "FourOverNotSeeGiveUp":0,
  "FourOverNotSeeFollow":90,
  "FourOverNotSeeRaise":5,
  "FourOverNotSeeCompare":5,
  "FourOverNotSeeAllin":0,

  "FourOverBelowKGiveUp":100,
  "FourOverDzGiveUp":10,
  "FourOverDzFollow":80,
  "FourOverDzRaise":10,
  "FourOverDzCompare":10,
  "FourOverDzAllin":0,

  "FourOverSzGiveUp":0,
  "FourOverSzAllin":5,
  "FourOverSzCompare":10,
  "FourOverSzRaise":30,
  "FourOverSzFollow":55,

  "FourOverJhGiveUp":0,
  "FourOverJhFollow":40,
  "FourOverJhRaise":50,
  "FourOverJhCompare":5,
  "FourOverJhAllin":5,

  "FourOverOverJhGiveUp":0,
  "FourOverOverJhFollow":40,
  "FourOverOverJhRaise":40,
  "FourOverOverJhCompare":0,
  "FourOverOverJhAllin":20,


  "FirstRoundRaise1": 60,
  "FirstRoundRaise2": 30,
  "FirstRoundRaise3": 10,
  "SecondRoundNotSeeRaise1": 60,
  "SecondRoundNotSeeRaise2": 30,
  "SecondRoundNotSeeRaise3": 10,
  "SecondRoundSawRaise1": 60,
  "SecondRoundSawRaise2": 30,
  "SecondRoundSawRaise3": 10,
  "ThirdRoundNotSeeRaise1": 60,
  "ThirdRoundNotSeeRaise2": 30,
  "ThirdRoundNotSeeRaise3": 10,
  "ThirdRoundDzRaise1": 60,
  "ThirdRoundDzRaise2": 30,
  "ThirdRoundDzRaise3": 10,
  "ThirdRoundSzRaise1": 10,
  "ThirdRoundSzRaise2": 60,
  "ThirdRoundSzRaise3": 30,
  "ThirdRoundJhRaise1": 10,
  "ThirdRoundJhRaise2": 30,
  "ThirdRoundJhRaise3": 60,
  "ThirdRoundOverJhRaise1": 10,
  "ThirdRoundOverJhRaise2": 30,
  "ThirdRoundOverJhRaise3": 60,
  "FourRoundNotSeeRaise1": 60,
  "FourRoundNotSeeRaise2": 30,
  "FourRoundNotSeeRaise3": 10,
  "FourRoundDzRaise1": 60,
  "FourRoundDzRaise2": 30,
  "FourRoundDzRaise3": 10,
  "FourRoundSzRaise1": 10,
  "FourRoundSzRaise2": 60,
  "FourRoundSzRaise3": 30,
  "FourRoundJhRaise1": 10,
  "FourRoundJhRaise2": 30,
  "FourRoundJhRaise3": 60,
  "FourRoundOverJhRaise1": 10,
  "FourRoundOverJhRaise2": 30,
  "FourRoundOverJhRaise3": 60
}

]`