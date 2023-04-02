import http from "k6/http";
import { sleep, check } from "k6";

export let options = {
    vus: 1000, // virtual users
    duration: '5m', // 5 minutes
    rps: 1000 // requests per second
};

export default function () {
    let res = http.get("http://localhost:8080/search?keyword=yoshua+bengio");
    check(res, {
        'response code was 200': (res) => res.status === 200,
    });

    sleep(1);
}