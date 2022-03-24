
//! 枚举类型
//* TypeScript特有的类型
//* 枚举的值可以随意设置，默认从0开始递增
enum HTTPStatus {
    OK = 200,
    NOT_FOUND = 404,
    INTERNAL_SERVER_ERROR = 500,
}

function processHttpStatus2(s: HTTPStatus) {
    if (s === HTTPStatus.OK) {
        console.log('good response')
    }

    // I want to print 枚举类型的name
    console.log(HTTPStatus[s])
}

processHttpStatus2(HTTPStatus.INTERNAL_SERVER_ERROR)
