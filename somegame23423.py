def add(x, y):
    return x + y

def subtract(x, y):
    return x - y

def multiply(x, y):
    return x * y

def divide(x, y):
    if y == 0:
        return "错误: 除数不能为零！"
    return x / y

def calculator():
    print("欢迎使用简单计算器")
    print("选择运算：")
    print("1. 加法 (+)")
    print("2. 减法 (-)")
    print("3. 乘法 (*)")
    print("4. 除法 (/)")

    while True:
        choice = input("请输入你的选择 (1/2/3/4)：")

        if choice in ('1', '2', '3', '4'):
            try:
                num1 = float(input("请输入第一个数字："))
                num2 = float(input("请输入第二个数字："))
            except ValueError:
                print("输入无效，请输入数字！")
                continue

            if choice == '1':
                print(f"{num1} + {num2} = {add(num1, num2)}")
            elif choice == '2':
                print(f"{num1} - {num2} = {subtract(num1, num2)}")
            elif choice == '3':
                print(f"{num1} * {num2} = {multiply(num1, num2)}")
            elif choice == '4':
                print(f"{num1} / {num2} = {divide(num1, num2)}")

        else:
            print("无效的输入，请选择 1/2/3/4")

        next_calculation = input("是否继续计算？ (yes/no)：").lower()
        if next_calculation != 'yes':
            print("感谢使用计算器！再见！")
            break

if __name__ == "__main__":
    calculator()
