"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.useGPT = void 0;
function useGPT() {
    const send = async (message) => {
        try {
            const resp = await wx.cloud.callContainer({
                "config": {
                    "env": "prod-0g1idfbhf182d0c2"
                },
                "path": "/chat/text_completion",
                "header": {
                    "X-WX-SERVICE": "golang-1jmj"
                },
                "method": "POST",
                "data": {
                    "action": "inc",
                    "prompt": message
                }
            });
            console.log(resp.data);
        }
        catch (error) {
            console.error(error);
        }
        return message;
    };
    return { send };
}
exports.useGPT = useGPT;
