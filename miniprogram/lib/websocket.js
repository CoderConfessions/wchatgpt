"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.useWebsocket = void 0;
function useWebsocket(data) {
    const socketTask = wx.connectSocket({
        url: "",
    });
    return {
        send: () => {
            socketTask.onOpen(() => socketTask.send({ data: JSON.stringify(data) }));
        },
        handleMessage: (handler) => {
            socketTask.onMessage((res) => handler(res.data));
        },
        handleClose: (handler) => {
            socketTask.onClose((res) => handler(res));
        },
        handleError: (handler) => {
            socketTask.onError((res) => handler(res));
        },
    };
}
exports.useWebsocket = useWebsocket;
