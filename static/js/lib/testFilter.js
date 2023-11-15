const TestFilter = {
    search: (testList, queryString, callback) => {
        if (!queryString || queryString.length == 0) {
            return callback(testList)
        }

        callback(testList.filter(test => test.Title.toLowerCase().includes(queryString.toLowerCase()) || test.ID == parseInt(queryString)))
    },

    sort: (testList, field, callback, type) => {
        callback(testList.sort((a, b) => {
            switch (type) {
                case "string": {
                    const aStr = a[field].toUpperCase()
                    const bStr = b[field].toUpperCase()

                    if (aStr < bStr) {
                        return -1;
                    }

                    if (aStr > bStr) {
                        return 1;
                    }

                    return 0
                }
                case "date": {
                    return new Date(b[field]) - new Date(a[field])
                }
            }
        }))
    }
}