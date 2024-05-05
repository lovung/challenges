package contest393

// class Solution {
// public:
//     long long findKthSmallest(vector<int>& coins, int k) {
//         using ll = long long;
//       int n = coins.size();
//       vector<ll> mult(1 << n, 1);
//       for (int b = 0; b < (1 << n); ++b) {
//         for (int i = 0; i < n; ++i) {
//           if (b & (1 << i)) {
//             mult[b] = mult[b] / gcd(mult[b], coins[i]) * coins[i];
//           }
//         }
//       }
//       ll lo = 1, hi = 1e12;
//       while (lo < hi) {
//         ll mid = (lo + hi) / 2;
//         ll cnt = 0;
//         for (int b = 1; b < (1 << n); ++b) {
//           ll v = mid / mult[b];
//           if (__builtin_popcount(b) & 1) cnt += v;
//           else cnt -= v;
//         }
//         if (cnt < k) lo = mid + 1;
//         else hi = mid;
//       }
//       return lo;
//     }
// };

func findKthSmallest_3(coins []int, k int) int64 {
	n := len(coins)
	mult := make([]int64, (1 << n))
	for i := range mult {
		mult[i] = 1
	}
	for b := 0; b < (1 << n); b++ {
		for i := 0; i < n; i++ {
			if b&(1<<i) != 0 {
				mult[b] = mult[b] / int64(gcd(int(mult[b]), coins[i])) * int64(coins[i])
			}
		}
	}
	var lo, hi int64 = 0, 1e12
	for lo < hi {
		mid := (lo + hi) / 2
		cnt := int64(0)
		for b := int64(1); b < (1 << n); b++ {
			v := mid / mult[b]
			if popcount(b)&1 != 0 {
				cnt += v
			} else {
				cnt -= v
			}
		}
		if cnt < int64(k) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func popcount(i int64) int64 {
	cnt := int64(0)
	for i > 0 {
		cnt += i & 1
		i >>= 1
	}
	return cnt
}
