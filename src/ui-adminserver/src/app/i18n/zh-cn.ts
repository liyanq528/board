import { COMMENT_MARKER } from '@angular/core/src/render3/interfaces/i18n';

export const LANG_ZH_CN = {
  BUTTON: {
    OK: '确定',
    YES: '是',
    NO: '否',
    CONFIRM: '确定',
    CANCEL: '取消',
    DELETE: '删除',
    CREATE: '创建',
    NEXT: '下一步',
    BACK: '上一步',
    TEST: '测试',
  },
  DATAGRID: {
    ITEMS: '条记录',
    TOTAL: ' 共:'
  },
  GLOBAL_ALERT: {
    TITLE: '标题',
    HINT: '提示',
    ASK: '询问',
    DELETE: '删除',
    SIGN_IN: '登录',
    WARNING: '警告',
    ERROR_DETAIL: '错误详情'
  },
  ERROR: {
    METHOD_NOT_ALLOWED: '不允许调用的方法。',
    CONFLICT_INPUT: '存在一个与当前有冲突的资源。',
    INPUT_ONLY_NUMBER: '只能输入数字',
    INPUT_REQUIRED: '字段为必填项',
    INPUT_NOT_REPEAT: '字段不能重复',
    INPUT_MAX_LENGTH: '最大长度',
    INPUT_MIN_LENGTH: '最小长度',
    INPUT_PATTERN: '输入不符合规则',
    INPUT_MAX_VALUE: '最大值',
    INPUT_MIN_VALUE: '最小值',
    HTTP_504: '网关超时:504',
    HTTP_500: '内部错误:500',
    HTTP_502: '网关错误:502',
    HTTP_400: '错误请求:400',
    HTTP_401: 'token已过期,请重新登录:401',
    HTTP_403: '您没有足够的权限进行该操作:403',
    HTTP_404: '资源未找到:404',
    HTTP_412: '请求内容不满足应提供的前置条件:412',
    HTTP_422: '无法处理实体:422',
    HTTP_UNK: '未知错误',
    HTTP_TIME_OUT: 'Http请求超时',
    INSUFFICIENT_PRIVILEGE: '权限不足，无法操作.'
  },
  HEAD_NAV: {
    DASHBOARD: '仪表盘',
    CONFIGURATIONS: '配置',
    RESOURCES: '资源',
    LANG_EN_US: 'English',
    LANG_ZH_CN: '中文',
    LOGOUT: '登出',
  },
  DASHBOARD: {
    BUTTONS: {
      RESTART_BORAD: '重启Board',
      STOP_BORAD: '停止Board',
      DETAIL: '详情',
      RESTART_CONTAINER: '重启容器',
      STOP_CONTAINER: '停止容器',
    },
    TITLES: {
      SYSTEM_INFO: '系统信息:',
      SYSTEM_CONTENT: '空',
      CONTAINER_INFO: '容器信息:',
      CONTAINER_CONTENT: '个容器正在运行',
    },
    CONTAINER_INFO: {
      NAME: '容器名: ',
      ID: 'ID: ',
      IMAGE: '镜像名:  ',
      CREATED: '创建时间: ',
      STATUS: '状态: ',
      PORTS: '端口号: ',
      CPU_RATE: 'CPU使用率: ',
      MEMORY_USAGE: '内存使用: ',
      MEMORY_RATE: '内存使用率: ',
      NETWORK_IO: '网络I/O: ',
      BLOCK_IO: '硬盘I/O: ',
      PID: 'PID: ',
    },
  },
  Resource: {
    Nav_Compute: '计算资源',
    Nav_Storage: '存储资源',
    Nav_Network: '网络资源',
    Nav_Node_List: '节点列表',
    Nav_Node_Log_List: '节点操作日志列表',
  },
  Node: {
    Node_List_Add: '添加节点',
    Node_List_Empty: '节点列表为空',
    Node_List_Remove_Node: '移除节点',
    Node_List_Remove_Ask: '确定要移除节点么？',
    Node_List_Status_Unknown: '未知',
    Node_List_Status_Schedulable: '可调度',
    Node_List_Status_Unschedulable: '不可调度',
    Node_Logs_Ip: 'IP',
    Node_Logs_Operation_Time: '操作时间',
    Node_Logs_PID: 'PID',
    Node_Logs_Type: '添加/移除',
    Node_Logs_Success: '是否成功',
    Node_Logs_Operation: '操作',
    Node_Logs_Empty: '节点操作日志为空',
    Node_Logs_Delete_Ask: '确定要删除这个日志?',
    Node_Logs_Delete_Log_Success: '删除日志成功',
    Node_Logs_Delete_Log_In_Use: '有节点在使用该日志，不能被删除。',
    Node_Logs_Stop_Ask: '确定要停止此操作么?',
    Node_Logs_Stop_Success: '已停止。',
    Node_Logs_Stop: '停止',
    Node_Logs_Removing: '正在移除...',
    Node_Logs_Adding: '正在添加...',
    Node_Detail_Add: '添加',
    Node_Detail_Remove: '移除',
    Node_Detail_Refresh: '刷新日志',
    Node_Detail_Title_Add: '添加节点',
    Node_Detail_Title_Remove: '正在移除节点',
    Node_Detail_Title_Log: '添加节点日志',
    Node_Detail_Node_Ip: '节点IP',
    Node_Detail_Node_Password: '节点密码',
    Node_Detail_Remove_Success: '移除节点成功。',
    Node_Detail_Add_Success: '添加节点成功。',
    Node_Detail_Remove_Failed: '移除节点失败。',
    Node_Detail_Add_Failed: '添加节点失败。',
    Node_Detail_Master_Title: 'Master {{0}} 密码',
    Node_Detail_Master_Hint: '请输入Master的密码',
    Node_Detail_Host_Username: `主机 {{0}} 用户名`,
    Node_Detail_Host_Password: `主机 {{0}} 密码`,
    Node_Detail_Host_Username_Hint: '请输入主机的用户名',
    Node_Detail_Host_Password_Hint: '请输入主机的密码',
    Node_Detail_Node_Ip_Hint: '请输入节点IP',
    Node_Detail_Node_Password_Hint: '请输入节点密码'
  },
  CONFIGURATIONS: {
    API_SERVER: {
      HOSTNAME: {
        NAME: 'Board IP地址',
        TIPS: 'Board服务部署的主机地址，不要使用localhost或者127.0.0.1上，因为Board需要被其他节点访问'
      },
      API_SERVER_PORT: {
        NAME: 'API服务器端口号',
        TIPS: 'Board服务部署的端口号'
      },
      KUBE_HTTP_SCHEME: {
        NAME: 'K8s网络请求方案',
        TIPS: 'Kubernetes(K8s)部署使用的网络协议，支持http和https'
      },
      KUBE_MASTER_IP: {
        NAME: 'K8s主节点IP',
        TIPS: 'Kubernetes(K8s)部署的主机地址，不要使用localhost或者127.0.0.1上，因为K8s需要被其他节点访问'
      },
      KUBE_MASTER_PORT: {
        NAME: 'K8s主节点端口号',
        TIPS: 'Kubernetes(K8s)部署的端口号，如果网络协议为http，则使用8080端口；如果使用https则使用6443端口'
      },
      REGISTRY_IP: {
        NAME: '镜像仓库IP',
        TIPS: '镜像仓库的地址'
      },
      REGISTRY_PORT: {
        NAME: '镜像仓库端口号',
        TIPS: '镜像仓库的端口号'
      },
      IMAGE_BASELINE_TIME: {
        NAME: '镜像基准时间',
        TIPS: '镜像的基准时间，默认为2016-01-01 09:00:00'
      },
    },
    GOGITS: {
      GOGITS_HOST_IP: {
        NAME: '服务IP地址',
        TIPS: '服务部署的主机地址，不要使用localhost或者127.0.0.1上，因为需要被其他节点访问'
      },
      GOGITS_HOST_PORT: {
        NAME: '服务端口号',
        TIPS: '服务部署的端口号'
      },
      GOGITS_SSH_PORT: {
        NAME: 'SSH请求端口号',
        TIPS: 'SSH(远程登录会话协议)的请求端口号'
      },
    },
    JENKINS: {
      JENKINS_HOST_IP: {
        NAME: '主机IP地址',
        TIPS: 'Jenkins master节点的地址，用于外部访问'
      },
      JENKINS_HOST_PORT: {
        NAME: '主机端口号',
        TIPS: 'Jenkins master节点的端口号'
      },
      JENKINS_NODE_IP: {
        NAME: '节点IP地址',
        TIPS: 'Jenkins node节点的地址'
      },
      JENKINS_NODE_SSH_PORT: {
        NAME: '节点SSH请求端口号',
        TIPS: 'Node节点的SSH(远程登录会话协议)请求端口号'
      },
      JENKINS_NODE_USERNAME: {
        NAME: '节点用户名',
        TIPS: '登录Node主机的用户名'
      },
      JENKINS_NODE_PASSWORD: {
        NAME: '节点密码',
        TIPS: '登录Node节点的密码'
      },
      JENKINS_NODE_VOLUME: {
        NAME: '节点数据卷路径',
        TIPS: 'Node节点的数据卷的地址'
      },
      JENKINS_EXECUTION_MODE: {
        NAME: '运行模式',
        TIPS: 'Jenkins集群的运行模式，支持单节点(single)或者多节点(multi)'
      },
    },
    KVM: {
      KVM_REGISTRY_SIZE: {
        NAME: 'KVM仓库大小',
        TIPS: '',
      },
      KVM_REGISTRY_PORT: {
        NAME: 'KVM仓库端口号',
        TIPS: '',
      },
      KVM_REGISTRY_PATH: {
        NAME: 'KVM工具路径',
        TIPS: '',
      },
    },
    LDAP: {
      LDAP_URL: {
        NAME: '服务URL地址',
        TIPS: 'LDAP服务的入口地址',
      },
      LDAP_BASEDN: {
        NAME: '服务DN来源',
        TIPS: 'LDAP验证用户是来自哪个源，ou为people,dc为mydomain,dc为com',
      },
      LDAP_UID: {
        NAME: '服务匹配源',
        TIPS: '搜索中用于匹配用户的属性，可以是uid，cn，email，sAMAccountName或其他属性，具体取决于您的LDAP/AD',
      },
      LDAP_SCOPE: {
        NAME: '服务匹配范围',
        TIPS: '搜索用户的范围（LDAP_SCOPE_BASE，LDAP_SCOPE_ONELEVEL，LDAP_SCOPE_SUBTREE）',
      },
      LDAP_TIMEOUT: {
        NAME: '请求超时时间',
        TIPS: '连接到LDAP服务器的超时时间（以秒为单位）。默认值（最合理）是5秒',
      },
    },
    EMAIL: {
      EMAIL_IDENTITY: {
        NAME: '身份',
        TIPS: '身份留为空白，用作用户名',
      },
      EMAIL_SERVER: {
        NAME: '服务IP地址',
        TIPS: '邮箱服务器地址',
      },
      EMAIL_SERVER_PORT: {
        NAME: '服务端口号',
        TIPS: '邮箱服务器的端口号',
      },
      EMAIL_USERNAME: {
        NAME: '用户名',
        TIPS: '',
      },
      EMAIL_PASSWORD: {
        NAME: '邮箱密码',
        TIPS: '',
      },
      EMAIL_FROM: {
        NAME: '邮件来源',
        TIPS: '',
      },
      EMAIL_SSL: {
        NAME: '是否启用SSL',
        TIPS: '',
      },
    },
    OTHERS: {
      ARCH_TYPE: {
        NAME: '系统架构',
        TIPS: '默认架构是x86_64，现在也支持mips',
      },
      DATABASE_PASSWORD: {
        NAME: '数据库密码',
        TIPS: 'mysql db的root用户的密码，请在生产前使用此密码',
      },
      TOKEN_CACHE_EXPIRE: {
        NAME: 'Token缓存存留时间',
        TIPS: '存储在缓存中的令牌的到期秒数',
      },
      TOKEN_EXPIRE: {
        NAME: 'Token存留时间',
        TIPS: '令牌的到期秒数',
      },
      ELASETICSEARCH_MEMORY: {
        NAME: 'Elasticsearch内存',
        TIPS: 'Elasticsearch可以使用的最大内存（MB）',
      },
      TILLER_PORT: {
        NAME: 'Tiller端口号',
        TIPS: 'Helm tiller的节点端口号',
      },
      BOARD_ADMIN_PASSWORD: {
        NAME: 'Board Admin密码',
        TIPS: 'Board Admin的初始密码仅在Board首次启动时才能生效，请在Board启动成功后再进行修改。',
      },
      BOARD_ADMIN_PASSWORD_OLD: {
        NAME: 'Board Admin旧密码',
        TIPS: '请输入Board Admin旧密码用于系统验证，验证通过则可以在下方修改新密码。',
      },
      BOARD_ADMIN_PASSWORD_NEW: {
        NAME: '新密码',
        TIPS: '请输入Board Admin新密码。',
      },
      BOARD_ADMIN_PASSWORD_CONFIRM: {
        NAME: '确认密码',
        TIPS: '请确认一次上面输入的新密码。',
      },
      AUTH_MODE: {
        NAME: '用户验证模式',
        TIPS: '默认情况下，身份验证模式为db_auth，即凭据存储在本地数据库中；如果要针对LDAP服务器验证用户的凭据，请将其设置为ldap_auth；如果要根据InData集成平台验证用户的凭据，请将其设置为indata_auth',
      },
      VERIFICATION_URL: {
        NAME: '验证地址',
        TIPS: '外部令牌验证URL，用于与其他平台集成授权。注：仅当auth_mode设置为\'indata_auth\'时，此选项才可用',
      },
      REDIRECTION_URL: {
        NAME: '重定向地址',
        TIPS: '当令牌无效或UI将重定向到的过期时，请指定重定向URL。注：仅当auth_mode设置为\'indata_auth\'时，此选项才可用',
      },
      AUDIT_DEBUG: {
        NAME: '审计',
        TIPS: '记录运行审核功能的所有运行切换',
      },
      DNS_SUFFIX: {
        NAME: 'DNS后缀',
        TIPS: 'Kubernetes DNS后缀',
      },
    },
  },
  ACCOUNT: {
    SIGN_IN: '登 录',
    USERNAME_PLACEHOLDER: '用户名',
    PASSWORD_PLACEHOLDER: '密码',
    SIGN_UP: '立即注册',
    FORGOT_PASSWORD: '忘记密码',
    FORGOT_PASSWORD_HINT_MSG: '请输入用户名/邮箱',
    FORGOT_PASSWORD_ERROR_MSG: '用户名/邮箱不能为空',
    USERNAME: '用户名',
    REALNAME: '真实姓名',
    INPUT_USERNAME: '输入用户名',
    EMAIL: '邮箱',
    INPUT_EMAIL: '输入邮箱',
    PASSWORD: '密码',
    INPUT_PASSWORD: '输入密码',
    CONFIRM_PASSWORD: '确认密码',
    INPUT_CONFIRM_PASSWORD: '输入确认密码',
    COMMENT: '备注',
    REQUIRED_ITEMS: '为必填项。',
    BACK: '返回',
    REGISTER: '用户注册',
    USERNAME_IS_REQUIRED: '用户名为必填项。',
    USERNAME_ALREADY_EXISTS: '用户名已存在。',
    USERNAME_IS_KEY: '不能用关键字作为用户名。',
    USERNAME_ARE_NOT_IDENTICAL: '用户名由数字、小写字母以及下划线组成，且长度范围是[4,40]。',
    EMAIL_IS_REQUIRED: '邮箱为必填项。',
    EMAIL_IS_ILLEGAL: '邮箱格式非法。',
    EMAIL_ALREADY_EXISTS: '邮箱已存在。',
    PASSWORD_IS_REQUIRED: '密码为必填项。',
    PASSWORDS_ARE_NOT_IDENTICAL: '两次密码内容不一致。',
    PASSWORD_FORMAT: '密码长度至少为8且需包含至少一个大写字符，一个小写字符和一个数字。',
    ERROR: '错误',
    INCORRECT_USERNAME_OR_PASSWORD: '用户名或密码错误。',
    SUCCESS_TO_SIGN_IN: '用户登录成功',
    FAILED_TO_SIGN_IN: '登录失败:',
    FAILED_TO_SIGN_UP: '注册失败:',
    SUCCESS_TO_SIGN_UP: '注册用户成功。',
    FAILED_TO_SIGN_OUT: '登出系统失败。',
    ACCOUNT_SETTING_SUCCESS: '账户设置成功。',
    ALREADY_SIGNED_IN: '该用户已在其他地方登录。',
    SEND_REQUEST: '发送验证',
    SEND_REQUEST_SUCCESS: '发送验证成功',
    SEND_REQUEST_SUCCESS_MSG: '发送验证成功，请检查邮件，访问邮件中的重置密码',
    RESET_NEW_PASS: '设置新密码',
    USER_NOT_EXISTS: '用户不存在。',
    SEND_REQUEST_ERR: '发送验证请求失败',
    SEND_REQUEST_ERR_MSG: '发送验证请求失败，请联系管理员。',
    RESET_PASS_SUCCESS: '重置密码成功',
    RESET_PASS_SUCCESS_MSG: '重置密码成功，请重新登录。',
    RESET_PASS_ERR: '重置密码失败',
    RESET_PASS_ERR_MSG: '重置密码失败，请联系管理员。',
    INVALID_RESET_UUID: '链接失效。',
    REGISTER_INFO: '系统刚刚初始化，您需要先设置一个帐户和密码，然后才能使用它。',
  },
  CONFIGURATIONPAGE: {
    UPLOAD: '上传',
    SEARCH: '搜索配置',
    CURRENT: '当前配置',
    TEMPORARY: '未生效配置',
    CURRENTTIP: '这里显示了当前Board的配置。您可以单击旁边的按钮以获取未生效的配置。',
    TEMPORARYTIP: '这里显示了未生效的配置。您可以单击旁边的按钮以查看Board当前配置。',
    SAVETOSERVER: '保存到服务',
    SAVETOLOCAL: '下载到本地',
    SAVECONFIGURATION: {
      TITLE: '是否生效配置？',
      SAVE: '保存成功',
      COMMENT: '新配置已经成功保存，是否在Board中应用该配置？生效配置需要花费一些时间来重启系统，确认生效请输入该主机的管理员账户及密码。',
      ACCOUNT: {
        NAME: '账户',
        PASSWORD: '密码',
        REQUIRED: '必填项'
      },
      CANCEL: '不生效',
      APPLY: '生效配置',
    }
  },
  INITIALIZATION: {
    TITLE: 'Adminserver初始化',
    BUTTONS: {
      NEXT: '下一步',
      TRANSLATE: 'Translate to English',
      APPLY: '应用',
      RESTART_DB: '重启数据库',
    },
    PAGE_TITLE: {
      WELCOME: '欢迎使用Adminserver',
      UUID: '验证UUID',
      DATABASE: '配置数据库',
      SSH: '启动数据库',
      ACCOUNT: '初始化账户',
    },
    PAGE_NAV_TITLE: {
      WELCOME: '欢迎页',
      UUID: '验证UUID',
      DATABASE: '配置数据库',
      SSH: '启动数据库',
      ACCOUNT: '初始化账户',
    },
    TOOLS: {
      LOADING: '等待中···',
    },
    ALERTS: {
      INITIALIZATION: '初始化错误！请检查Docker是否正常运行。',
      UUID: 'UUID错误！请重新输入！',
      DATABASE: '数据库初始化失败！请检查Docker是否正常运行。',
      PASSWORDS_IDENTICAL: '两次密码不一致！',
      MAX_CONNECTION: '数据库最大连接数错误！',
      SSH: '网络错误或者账户错误，请重试！',
      ACCOUNT: '网络错误或者docker运行错误，请重试！',
    },
    CONTENTS: {
      WELCOME: '欢迎使用Adminserver！由于这是第一次进行初始化流程，因此需要完成以下几项配置后才能正常启动系统。',
      UUID: '为了确认您的身份，需要您输入/data/board/secrets 文件夹中的UUID以进行下一步的配置。',
      DATABASE: '您需要配置数据库密码以初始化数据库。该步骤可能需要一些时间。',
      SSH: '请输入当前主机的账户及密码。注意：账户需要一定的权限用于安装并运行相关的组件。运行需要一些时间。系统不会存储您的账户及密码。',
      ACCOUNT: '初始化adminserver的管理员密码。该admin账户将与board共用。',
    },
    LABELS: {
      UUID: 'UUID',
      MAX_CONNECTION: '最大连接数',
      PASSWORD: '密码',
      PASSWORD_CONFIRM: '确认密码',
      ACCOUNT: '账户',
    },
    HELPER: {
      REQUIRED: '必填项！',
      NUMS: '接受10~16384。',
      MIN_LENGTH_8: '至少包含8个字符！',
      NOT_EDITABLE: '不允许更改！',
    },
    MODAL_TITLE: '数据库错误！',
  },
};
