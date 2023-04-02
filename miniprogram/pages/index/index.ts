import { useConversation } from "./conversation";
import { useGPT } from "./gpt";

Page({
  data: {
    $conversation: [],
    showStartPage: true,
    isGenerating: false,
    prompt: "",
    keyboardHeight: "0px",
  },

  onLoad() {
    wx.cloud.init();
  },

  onShareAppMessage() {
    return {
      title: "邀您体验免费的 GPT-3.5",
      path: "/pages/index/index",
    };
  },

  onShareTimeline() {
    return {
      title: "邀您体验免费的 GPT-3.5",
    };
  },

  onKeyboardPullUp(event: any) {
    this.setData({ keyboardHeight: event.detail.height + "px" });
  },

  onKeyboardDismiss() {
    this.setData({ keyboardHeight: "0px" });
  },

  onTapExample(event: any) {
    this.setData({ prompt: event.currentTarget.dataset.prompt });
  },

  onTapMessage(event: any) {
    const i = event.currentTarget.dataset.i;
    const content = this.data.$conversation[i][1];

    wx.setClipboardData({ data: content + "\n\n——微信小程序 wChatGPT" });
  },

  async onTapSendButton() {
    this.setData({ isGenerating: true });

    const conversation = useConversation(this);
    const prompt = this.data.prompt.trim();
    if (prompt === "") {
      wx.showToast({ title: "请输入内容", icon: "none" });
      return;
    }

    conversation.addUserMessage(prompt);
    conversation.addAssistantMessage("");

    this.setData({ showStartPage: false, prompt: "" });

    const history = this.data.$conversation.slice(0, this.data.$conversation.length - 1);
    const { send } = useGPT();
    const final = () => this.setData({ isGenerating: false });

		const serializedMessages = JSON.stringify(history)
		console.log(serializedMessages);
		send(prompt).then((response: any) => {
			// 将conversation（类型为数组）中最后一个
			conversation.appendAssistantMessage(prompt);
			final();
		}).catch(() => {
			conversation.appendAssistantMessage("网络错误");
			final();
		});
  },
});
