import * as fs from 'fs';
import * as path from 'path';
import * as mammoth from 'mammoth';

/**npm install mammoth
 npm install @types/mammoth

 * 读取指定目录下的所有 Word 文件，并返回其内容
 * @param dirPath 要读取的目录路径
 * @returns 返回所有 Word 文件的内容
 */
async function readWordFilesFromDir(dirPath: string): Promise<void> {
    const files = await fs.promises.readdir(dirPath);  // 读取目录下的所有文件
    const wordFiles = files.filter(file => file.endsWith('.docx'));  // 过滤出所有 .docx 文件


    const contentPromises: Promise<string>[] = wordFiles.map(async file => {
        const filePath = path.join(dirPath, file);  // 拼接文件完整路径
        await readWordFileContent(filePath);  // 读取文件内容
    });

    // 等待所有文件的内容读取完毕
  //  const contents = await Promise.all(contentPromises);
  //  return contents;
}

/**
 * 使用 mammoth 库读取 Word 文件内容
 * @param filePath 要读取的 Word 文件路径
 * @returns 返回文件的纯文本内容
 */
function readWordFileContent(filePath: string): Promise<string> {
    return new Promise((resolve, reject) => {
        const stream = fs.createReadStream(filePath);  // 创建文件读取流

        mammoth.extractRawText({ stream: stream })  // 使用 mammoth 提取文件的纯文本
            .then(result => {
                resolve(result.value);  // 返回文件的纯文本内容
            })
            .catch(err => {
                reject(`无法读取文件 ${filePath}: ${err.message}`);
            });
    });
}

// 使用示例
async function main() {
    const dirPath = 'C:\\Users\\attil\\OneDrive\\sumdoc2023 git';  // 替换为你的目录路径


        await readWordFilesFromDir(dirPath);

}

main();
