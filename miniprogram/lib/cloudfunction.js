"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.useMsgSecCheck = void 0;
async function useMsgSecCheck(content) {
    const res = await wx.cloud.callFunction({
        name: "msgSecCheck",
        data: { content },
    });
    return { pass: res.result.result.suggest === "pass" };
}
exports.useMsgSecCheck = useMsgSecCheck;
