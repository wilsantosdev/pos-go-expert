import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    {
      duration: '10s',
      target: 1,
      exec: 'bothScenarios',
    },
    {
      duration: '10s',
      target: 11,
      exec: 'bothScenarios',
    },
    {
      duration: '10s',
      target: 101,
      exec: 'bothScenarios',
    }
  ]
};

export default function bothScenarios() {
  withoutToken();
  withToken();
}

export function withoutToken() {
  http.get('http://app:8000/v1/status');
  sleep(1);
}

export function withToken() {
  http.get('http://app:8000/v1/status', { headers: { API_KEY: '1234567890' } });
  sleep(1);
}
