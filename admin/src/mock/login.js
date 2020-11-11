const USER = {
    name: 'admin',
    user_id: 1,
    token: 'admin',
    avatar: 'https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png',
    router: [
        {
            path: 'web-settings',
            name: 'WebSetting',
            meta: {
                title: '网站设置',
                icon: 'global'
            }
        },
        {
            path: 'label',
            name: 'LabelSetting',
            meta: {
                title: '标签设置',
                icon: 'tags'
            }
        },
        {
            path: 'article',
            name: 'ArticleList',
            meta: {
                title: '文章列表',
                icon: 'file-markdown'
            }
        },
    ]
}

export const login = req => {
    req = JSON.parse(req.body)
    return { token: USER[req.userName].token}
}