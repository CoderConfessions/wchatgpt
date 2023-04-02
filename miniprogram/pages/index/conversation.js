"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.useConversation = void 0;
class Conversation {
    constructor(page) {
        this.page = page;
    }
    addUserMessage(prompt) {
        this.addMessage("user", prompt);
    }
    addAssistantMessage(content) {
        this.addMessage("assistant", content);
    }
    appendAssistantMessage(piece) {
        this.updateLatestAssistantMessage((msg) => {
            msg[1] += piece;
            msg[2].isPending = false;
        });
    }
    setAssistantMessageFailed() {
        this.updateLatestAssistantMessage((msg) => {
            msg[2].isPending = false;
            msg[2].isFailed = true;
        });
    }
    updateLatestAssistantMessage(modifier) {
        const i = this.page.data.$conversation.length - 1;
        const msg = this.page.data.$conversation[i];
        modifier(msg);
        const key = `$conversation[${i}]`;
        this.page.setData({ [key]: msg });
    }
    addMessage(role, content) {
        this.page.data.$conversation.push([role, content.trim(), { isPending: !content }]);
        this.page.setData({
            $conversation: this.page.data.$conversation,
        });
    }
}
let conversation;
function useConversation(page) {
    if (!conversation) {
        conversation = new Conversation(page);
    }
    return conversation;
}
exports.useConversation = useConversation;
