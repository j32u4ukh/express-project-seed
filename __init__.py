import argparse
import os
import shutil

# 檢查指定資料夾是否為空
def is_empty_folder(path) -> bool:
    # 取得路徑中所有檔案和資料夾，如果為空，則返回 True
    return os.listdir(path) is None

def copy_project(output: str, name: str):
    folder = os.path.join(output, name)
    if not os.path.exists(folder):
        os.mkdir(folder)

    shutil.copytree(src="./seed/public", dst=os.path.join(folder, "public"))
    shutil.copytree(src="./seed/views", dst=os.path.join(folder, "views"))
    shutil.copyfile(src="./seed/index.js", dst=os.path.join(folder, "index.js"))
    shutil.copyfile(src="./seed/package.json", dst=os.path.join(folder, "package.json"))
    shutil.copyfile(src="./seed/package-lock.json", dst=os.path.join(folder, "package-lock.json"))

def replace_content(path: str, origin_content: str, new_content: str) -> str:
    try:
        with open(path, 'r', encoding="utf-8") as file:
            content = file.read()

        new_content = content.replace(origin_content, new_content)
        
        with open(path, 'w', encoding="utf-8") as file:
            file.write(new_content)

        return ""
    except Exception as e:
        return str(e)
    
def main() -> None:
    parser = argparse.ArgumentParser(description="""初始化的 Express 專案 生成工具""")
    
    parser.add_argument('--output', '-o', required=True, help='專案輸出的資料夾')
    parser.add_argument('--name', '-n', required=True, help='專案名稱')
    parser.add_argument('--description', '-d', help='package.json 當中的 description 內容')

    # 解析 Command line 參數
    args = parser.parse_args()
    output = args.output

    if not os.path.exists(output):
        print("--output/-o 路徑不存在")
        return

    # 檢查路徑是否為資料夾
    if not os.path.isdir(output):
        print("--output/-o 參數必須輸入資料夾路徑")
        return
 
    path = os.path.join(output, "seed/views/layouts/main.hbs")
    name = args.name
    result = replace_content(path=path,
                             origin_content="<title>Project Title</title>", 
                             new_content=f"<title>{name}</title>")

    if result != "":
        print(result)
        return

    if args.description is not None:
        description = args.description
    else:
        description = "No description"
        
if __name__ == '__main__':
    main()
