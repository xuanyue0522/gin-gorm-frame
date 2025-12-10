module.exports = {
    messages: {
        type: '选择本次提交类型：',
        customScope: '\n请输入修改范围（可选）：',
        subject: '填写简短精炼的变更描述：\n',
        body: '填写更详细的变更描述（可选），使用 "|" 换行：\n',
        confirmCommit: '确认使用以上信息提交？(y/n)'
    },
    skipQuestions: ['footer'],
    types: [
        { value: 'feat', name: 'feat:     添加新功能' },
        { value: 'fix', name: 'fix:      修复Bug' },
        { value: 'docs', name: 'docs:     文档变更' },
        { value: 'style', name: 'style:    代码格式（不影响代码运行的变动）' },
        { value: 'refactor', name: 'refactor: 代码重构（既不是新增功能，也不是Bug修复）' },
        { value: 'perf', name: 'perf:     性能优化' },
        { value: 'test', name: 'test:     添加缺失的测试或修改现有的测试' },
        { value: 'chore', name: 'chore:    构建过程或辅助工具的变动' },
        { value: 'revert', name: 'revert:   回退提交' },
        { value: 'build', name: 'build:    打包编译相关变动' }
    ],
    breaklineChar: '|',
    subjectLimit: 100,
}