// index.ts
// 获取应用实例
const app = getApp<IAppOption>()

Page({
  data: {
    messages: [
      {"isMe": 1, "content": "my first message!"},
      {"isMe": 0, "content": "我是没有感情的gpt，请正确描述你的问题，我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题"},
      {"isMe": 1, "content": "my second message!my second message!my second message!my second message!my second message!my second message!my second message!my second message!my second message!my second message!my second message!my second message!my second message!"},
      {"isMe": 0, "content": "我是没有感情的gpt，请正确描述你的问题，我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题我是没有感情的gpt，请正确描述你的问题"}
    ],
    toView: "",
    userInfo: {},
    userInputValue: "", // 如需尝试获取用户信息可改为false
  },
  onLoad() {
    // @ts-ignore
    if (wx.getUserProfile) {
      this.setData({
        canIUseGetUserProfile: true
      })
    }
  },
  getUserProfile() {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log(res)
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    })
  },
  userInput: function (userValue:JSON) {
    this.setData({
      userInputValue: userValue.detail.value
    })
    userValue.detail.value = ""
  },
  onSend: function () {
    console.log(this.data.messages.length)
    this.userInput
    if (this.data.userInputValue.length == 0) {
      console.log("输入为空，不执行操作")
      return
    }
    this.data.messages.push({"isMe": 1, "content": this.data.userInputValue})
    this.data.messages.push({"isMe": 0, "content": this.data.userInputValue})
    console.log('msg-' + (this.data.messages.length - 1));
    
    this.setData({
      messages: this.data.messages,
      toView: 'msg-' + (this.data.messages.length - 1),
      userInputValue: ""
    })
  },
})
