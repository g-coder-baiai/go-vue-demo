module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/essential',
    '@vue/airbnb',
  ],
  parserOptions: {
    parser: 'babel-eslint',
  },
  rules: {
    // indent 项目的缩进 第一个参数表示报错级别（error.. warn.. off表示关闭) ,第二个参数表示缩进的格数
    'indent':['off'],
    // semi 表示是否使用分号,第一个参数表示报错级别（error.. warn.. off表示关闭) always never 
    // 'semi':['error','always'],
    // quotes表示使用单引号还是双引号
    // quotes:['error','double/single'];
    'arrow-body-style':['off'],
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
  },
};
