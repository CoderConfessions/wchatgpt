.PHONY: clean

# 编译 TypeScript 代码
compile:
    cd cloudfunctions && tsc
    cd miniprogram && tsc

# 编译 LESS 样式文件
style:
    cd miniprogram && lessc style.less style.css

# 运行项目
run:
    cd miniprogram && wechatdevtools .

# 清理编译产生的文件
clean:
    cd cloudfunctions && rm -rf dist
    cd miniprogram && rm -rf dist
    rm -f miniprogram/style.css