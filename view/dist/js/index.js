/**
 * Display the current time
 */
function getTime() {
    var dt = Date();
    console.log(dt.toString());
    return document.getElementById("clock_area").innerHTML = dt;
}
setInterval('getTime()', 1000);
