import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
      {
          duration: '10s',
          target: 1
      },
      {
          duration: '10s',
          target: 10
      },
      {
          duration: '10s',
          target: 100
      }
  ]
};

export default function() {
  http.get('http://app:8000/v1/status');
  sleep(1);
}
