/**
 * @param {number} rowsCount
 * @param {number} colsCount
 * @return {Array<Array<number>>}
 */
Array.prototype.snail = function (rowsCount, colsCount) {
    if (rowsCount * colsCount !== this.length) {
        return []
    }
    if (rowsCount === 1) {
        return this
    }
    let arr = []
    let idx = 0
    for (let i = 0; i < colsCount; i++) {
        for (let j = 0; j < rowsCount; j++) {
            arr[j] ??= []
            const val = i % 2 === 0 ? j : colsCount === rowsCount ? colsCount - j - 1 : rowsCount - j - 1
            arr[val].push(this[idx])
            idx++
        }
    }
    console.log(arr)
}

const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35]
arr.snail(7, 5)