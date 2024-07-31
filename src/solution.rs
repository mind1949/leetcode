/// [Leetcode](https://leetcode.cn) solution
pub struct Solution;

use std::collections::{HashMap, VecDeque};

impl Solution {
    /// [1. Tow Sum](https://leetcode.cn/problems/two-sum/description)
    ///
    /// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
    ///
    /// You may assume that each input would have exactly one solution, and you may not use the same element twice.
    ///
    /// You can return the answer in any order.
    ///
    ///  
    ///
    /// Example 1:
    ///
    //      Input: nums = [2,7,11,15], target = 9
    ///     Output: [0,1]
    ///     Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
    ///
    /// Example 2:
    ///
    ///     Input: nums = [3,2,4], target = 6
    ///     Output: [1,2]
    ///
    /// Example 3:
    ///
    ///     Input: nums = [3,3], target = 6
    ///     Output: [0,1]
    ///     
    ///  
    ///
    /// Constraints:
    ///
    ///     2 <= nums.length <= 104
    ///     -109 <= nums[i] <= 109
    ///     -109 <= target <= 109
    ///     Only one valid answer exists.
    ///
    ///  
    /// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut indices = HashMap::new();
        for (i, v) in nums.iter().enumerate() {
            if let Some(&idx) = indices.get(&(target - v)) {
                return vec![i as i32, idx as i32];
            };
            indices.insert(v, i);
        }
        return vec![];
    }

    /// [49. Group Anagrams](https://leetcode.cn/problems/group-anagrams/description/)
    ///
    /// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
    ///
    /// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: strs = ["eat","tea","tan","ate","nat","bat"]
    /// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
    ///
    /// Example 2:
    ///
    /// Input: strs = [""]
    /// Output: [[""]]
    ///
    /// Example 3:
    ///
    /// Input: strs = ["a"]
    /// Output: [["a"]]
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     1 <= strs.length <= 104
    ///     0 <= strs[i].length <= 100
    ///     strs[i] consists of lowercase English letters.
    ///
    ///
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut matchers = HashMap::new();
        for str in strs.iter() {
            let mut matcher = [0; 26];
            for e in str.chars() {
                matcher[e as usize - 'a' as usize] += 1;
            }
            let group = matchers.entry(matcher).or_insert(vec![]);
            group.push(str.to_string());
        }

        let mut result: Vec<Vec<String>> = Vec::with_capacity(matchers.len());
        for (_, group) in matchers {
            result.push(group);
        }
        return result;
    }

    /// [128. Longest Consecutive Sequence](https://leetcode.cn/problems/longest-consecutive-sequence/description/)
    ///
    /// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
    ///
    /// You must write an algorithm that runs in O(n) time.
    ///
    /// Example 1:
    ///
    /// Input: nums = [100,4,200,1,3,2]
    /// Output: 4
    /// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
    ///
    /// Example 2:
    ///
    /// Input: nums = [0,3,7,2,5,8,4,6,0,1]
    /// Output: 9
    ///
    /// Constraints:
    ///
    ///	0 <= nums.length <= 10^5
    ///	-10^9 <= nums[i] <= 10^9
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let mut exist: HashMap<i32, bool> = HashMap::new();
        for num in nums {
            exist.insert(num, true);
        }

        let mut longest = 0;
        for (&num, _) in exist.iter() {
            if let Some(&ok) = exist.get(&(num - 1)) {
                if ok {
                    continue;
                }
            }

            let mut length = 1;
            let mut n = num;
            while let Some(&ok) = exist.get(&(n + 1)) {
                if !ok {
                    break;
                }
                length += 1;
                n += 1;
            }
            longest = std::cmp::max(longest, length)
        }
        return longest;
    }

    /// [283. Move Zeroes](https://leetcode.cn/problems/move-zeroes/description/)
    ///
    /// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
    ///
    /// Note that you must do this in-place without making a copy of the array.
    ///
    /// Example 1:
    ///
    /// Input: nums = [0,1,0,3,12]
    /// Output: [1,3,12,0,0]
    ///
    /// Example 2:
    ///
    /// Input: nums = [0]
    /// Output: [0]
    ///
    /// Constraints:
    ///
    ///	1 <= nums.length <= 10^4
    ///	-2^31 <= nums[i] <= 2^31 - 1
    ///
    /// Follow up: Could you minimize the total number of operations done?
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut tail = 0;
        for head in 0..nums.len() {
            if nums[head] != 0 {
                nums.swap(tail, head);
                tail += 1;
            }
        }
    }

    /// [11. Container with most water](https://leetcode.cn/problems/container-with-most-water/description/)
    ///
    /// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
    ///
    /// Find two lines that together with the x-axis form a container, such that the container contains the most water.
    ///
    /// Return the maximum amount of water a container can store.
    ///
    /// Notice that you may not slant the container.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// ```txt
    ///  8|                    |                |              
    ///  7|                    |                |              |
    ///  6|          |         |          |     |              |
    ///  5|          |         |          |     |       |      |
    ///  4|          |         |          |     |       |      |
    ///  3|          |         |          |     |       |      |
    ///  2|          |         |    |     |     |       |      |
    ///  1|    |     |         |    |     |     |       |      |
    ///   -------------------------------------------------------
    ///   0    1     2         3    4     5     6       7      8
    /// ```
    ///
    /// Input: height = [1,8,6,2,5,4,8,3,7]
    ///
    /// Output: 49
    ///
    /// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
    ///
    /// Example 2:
    ///
    /// Input: height = [1,1]
    /// Output: 1
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     n == height.length
    ///     2 <= n <= 105
    ///     0 <= height[i] <= 104
    ///
    ///
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut i = 0;
        let mut j = height.len() - 1;
        let mut max_area = 0;

        while j > i {
            let area = height[i].min(height[j]) * (j - i) as i32;
            max_area = max_area.max(area);
            if height[i] > height[j] {
                j -= 1;
            } else {
                i += 1;
            }
        }
        return max_area;
    }

    /// [15. 3Sum](https://leetcode.cn/problems/3sum/description/)
    ///
    /// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
    ///
    /// Notice that the solution set must not contain duplicate triplets.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: nums = [-1,0,1,2,-1,-4]
    /// Output: [[-1,-1,2],[-1,0,1]]
    /// Explanation:
    /// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
    /// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
    /// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
    /// The distinct triplets are [-1,0,1] and [-1,-1,2].
    /// Notice that the order of the output and the order of the triplets does not matter.
    ///
    /// Example 2:
    ///
    /// Input: nums = [0,1,1]
    /// Output: []
    /// Explanation: The only possible triplet does not sum up to 0.
    ///
    /// Example 3:
    ///
    /// Input: nums = [0,0,0]
    /// Output: [[0,0,0]]
    /// Explanation: The only possible triplet sums up to 0.
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     3 <= nums.length <= 3000
    ///     -105 <= nums[i] <= 105
    ///     
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort();
        let mut result = Vec::new();
        for k in 0..(nums.len() - 2) {
            if nums[k] > 0 {
                break;
            }
            if k > 0 && nums[k] == nums[k - 1] {
                continue;
            }

            let mut i = k + 1;
            let mut j = nums.len() - 1;
            while i < j {
                let sum = nums[k] + nums[i] + nums[j];
                if sum == 0 {
                    result.push(vec![nums[k], nums[i], nums[j]])
                }
                if sum > 0 {
                    j = j - 1;
                    while i < j && nums[j] == nums[j + 1] {
                        j -= 1;
                    }
                } else {
                    i = i + 1;
                    while i < j && nums[i] == nums[i - 1] {
                        i += 1;
                    }
                }
            }
        }
        return result;
    }

    /// [42. Trapping Rain Water](https://leetcode.cn/problems/trapping-rain-water/description/)
    ///
    /// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
    /// Output: 6
    /// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
    ///
    /// Example 2:
    ///
    /// Input: height = [4,2,0,3,2,5]
    /// Output: 9
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     n == height.length
    ///     1 <= n <= 2 * 104
    ///     0 <= height[i] <= 105
    ///
    ///
    pub fn trap(height: Vec<i32>) -> i32 {
        if height.len() == 0 {
            return 0;
        }

        let mut total = 0;
        let mut left = 0;
        let mut right = height.len() - 1;
        let mut left_max = 0;
        let mut right_max = 0;

        while left < right {
            left_max = left_max.max(height[left]);
            right_max = right_max.max(height[right]);

            if height[left] < height[right] {
                total += left_max - height[left];
                left += 1;
            } else {
                total += right_max - height[right];
                right -= 1;
            }
        }
        return total;
    }

    /// [3. Longest Substring Without Repeating Characters](https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/)
    ///
    /// Given a string s, find the length of the longest
    /// substring
    /// without repeating characters.
    ///
    ///
    ///
    /// Example 1:
    ///
    /// Input: s = "abcabcbb"
    /// Output: 3
    /// Explanation: The answer is "abc", with the length of 3.
    ///
    /// Example 2:
    ///
    /// Input: s = "bbbbb"
    /// Output: 1
    /// Explanation: The answer is "b", with the length of 1.
    ///
    /// Example 3:
    ///
    /// Input: s = "pwwkew"
    /// Output: 3
    /// Explanation: The answer is "wke", with the length of 3.
    /// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
    ///
    ///
    ///
    /// Constraints:
    ///
    ///     0 <= s.length <= 5 * 104
    ///     s consists of English letters, digits, symbols and spaces.
    ///
    pub fn length_of_longest_substring(s: String) -> i32 {
        if s.len() <= 1 {
            return s.len() as i32;
        }

        let mut left = 0;
        let mut right = 0;
        let mut longest = 0;
        let mut m: HashMap<char, usize> = HashMap::new();
        let chars: Vec<char> = s.chars().collect();
        while left <= right && right < chars.len() {
            if let Some(&idx) = &m.get(&chars[right]) {
                for i in left..(idx) {
                    m.remove(&chars[i]);
                }
                left = idx + 1;
            } else {
                longest = longest.max(right - left + 1);
            }
            m.entry(chars[right]).or_insert(right);
            right += 1;
        }
        return longest as i32;
    }

    /// [438. Find All Anagrams in a String](https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/)
    ///
    /// Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.
    ///
    /// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: s = "cbaebabacd", p = "abc"
    /// Output: [0,6]
    /// Explanation:
    /// The substring with start index = 0 is "cba", which is an anagram of "abc".
    /// The substring with start index = 6 is "bac", which is an anagram of "abc".
    ///
    /// Example 2:
    ///
    /// Input: s = "abab", p = "ab"
    /// Output: [0,1,2]
    /// Explanation:
    /// The substring with start index = 0 is "ab", which is an anagram of "ab".
    /// The substring with start index = 1 is "ba", which is an anagram of "ab".
    /// The substring with start index = 2 is "ab", which is an anagram of "ab".
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     1 <= s.length, p.length <= 3 * 104
    ///     s and p consist of lowercase English letters.
    pub fn find_anagrams(s: String, p: String) -> Vec<i32> {
        if s.len() < p.len() {
            return vec![];
        }

        let chars_p: Vec<char> = p.chars().collect();
        let mut count_p = [0; 26];
        for &c in &chars_p {
            count_p[c as usize - 'a' as usize] += 1;
        }

        let chars: Vec<char> = s.chars().collect();
        let mut count_sub = [0; 26];
        for &c in chars[0..chars_p.len()].iter() {
            count_sub[c as usize - 'a' as usize] += 1;
        }

        let mut res = vec![];
        let (mut i, mut j) = (0, p.len() - 1);
        while i <= j && j < chars.len() {
            if count_p == count_sub {
                res.push(i as i32);
            }

            if let Some(&c) = chars.get(i) {
                count_sub[c as usize - 'a' as usize] -= 1;
            }
            i += 1;
            j += 1;
            if let Some(&c) = chars.get(j) {
                count_sub[c as usize - 'a' as usize] += 1;
            }
        }

        return res;
    }

    /// [560. Subarray Sum Equal K](https://leetcode.cn/problems/subarray-sum-equals-k/description/)
    ///
    /// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
    ///
    /// A subarray is a contiguous non-empty sequence of elements within an array.
    ///
    /// Example 1:
    ///
    /// Input: nums = [1,1,1], k = 2
    /// Output: 2
    ///
    /// Example 2:
    ///
    /// Input: nums = [1,2,3], k = 3
    /// Output: 2
    ///
    /// Constraints:
    ///
    ///	1 <= nums.length <= 2 * 10^4
    ///	-1000 <= nums[i] <= 1000
    ///	-10^7 <= k <= 10^7
    pub fn subarray_sum(nums: Vec<i32>, k: i32) -> i32 {
        let mut res = 0;
        let mut m = HashMap::from([(0, 1)]);
        let (mut sum, mut i) = (0, 0);
        while i < nums.len() {
            sum += nums[i];
            if let Some(count) = m.get(&(sum - k)) {
                res += count;
            }
            m.entry(sum).and_modify(|count| *count += 1).or_insert(1);
            i += 1;
        }
        return res;
    }

    /// [239. Sliding Window Maximum](https://leetcode.cn/problems/sliding-window-maximum/description/)

    ///
    /// You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.
    ///
    /// Return the max sliding window.
    ///
    /// Example 1:
    ///
    /// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
    /// Output: [3,3,5,5,6,7]
    /// Explanation:
    /// Window position                Max
    /// ---------------               -----
    /// [1  3  -1] -3  5  3  6  7       3
    ///
    ///	1 [3  -1  -3] 5  3  6  7       3
    ///	1  3 [-1  -3  5] 3  6  7       5
    ///	1  3  -1 [-3  5  3] 6  7       5
    ///	1  3  -1  -3 [5  3  6] 7       6
    ///	1  3  -1  -3  5 [3  6  7]      7
    ///
    /// Example 2:
    ///
    /// Input: nums = [1], k = 1
    /// Output: [1]
    ///
    /// Constraints:
    ///
    ///	1 <= nums.length <= 10^5
    ///	-10^4 <= nums[i] <= 10^4
    ///	1 <= k <= nums.length
    pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let k = k as usize;
        let mut top_idx: VecDeque<usize> = VecDeque::new();

        // 定义一个普通函数，用于执行单调栈的 push 操作
        fn push(nums: &Vec<i32>, top_idx: &mut VecDeque<usize>, i: usize) {
            while !top_idx.is_empty() && nums[i] >= nums[*top_idx.back().unwrap()] {
                top_idx.pop_back();
            }
            top_idx.push_back(i);
        }

        for i in 0..k {
            push(&nums, &mut top_idx, i);
        }

        let mut res = Vec::with_capacity(nums.len() - k + 1);
        res.push(nums[top_idx[0]]);

        for i in k..nums.len() {
            push(&nums, &mut top_idx, i);
            while !top_idx.is_empty() && *top_idx.front().unwrap() <= i - k {
                top_idx.pop_front();
            }
            res.push(nums[top_idx[0]]);
        }

        res
    }

    /// [76. Minimum Window Substring](https://leetcode.cn/problems/minimum-window-substring/description/)
    ///
    /// Given two strings s and t of lengths m and n respectively, return the minimum window
    /// substring
    /// of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
    ///
    /// The testcases will be generated such that the answer is unique.
    ///
    /// Example 1:
    ///
    /// Input: s = "ADOBECODEBANC", t = "ABC"
    /// Output: "BANC"
    /// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
    ///
    /// Example 2:
    ///
    /// Input: s = "a", t = "a"
    /// Output: "a"
    /// Explanation: The entire string s is the minimum window.
    ///
    /// Example 3:
    ///
    /// Input: s = "a", t = "aa"
    /// Output: ""
    /// Explanation: Both 'a's from t must be included in the window.
    /// Since the largest window of s only has one 'a', return empty string.
    ///
    /// Constraints:
    ///
    ///	m == s.length
    ///	n == t.length
    ///	1 <= m, n <= 105
    ///	s and t consist of uppercase and lowercase English letters.
    ///
    /// Follow up: Could you find an algorithm that runs in O(m + n) time?
    pub fn min_window(s: String, t: String) -> String {
        if t.len() > s.len() {
            return "".to_string();
        }

        fn get(chars: &Vec<char>, i: usize) -> usize {
            chars[i] as usize - 'A' as usize
        }
        // 统计 t 中每个字符的个数
        let mut counter = [0; 58];
        let chars: Vec<char> = t.chars().collect();
        for i in 0..chars.len() {
            counter[get(&chars, i)] += 1;
        }

        fn contain(src: &[i32], target: &[i32]) -> bool {
            for (i, &v) in target.iter().enumerate() {
                if src[i] >= v {
                    continue;
                }
                return false;
            }
            return true;
        }
        let chars: Vec<char> = s.chars().collect();
        let mut counter_s = [0; 58];
        let (mut l, mut r) = (0, 0);
        let (mut i, mut j) = (0, 1);

        counter_s[get(&chars, 0)] += 1;
        while i < j && j <= chars.len() {
            if contain(&counter_s, &counter) {
                if r - l == 0 || j - i < r - l {
                    (l, r) = (i, j)
                }

                counter_s[get(&chars, i)] -= 1;
                i += 1;
            } else {
                j += 1;
                if j <= chars.len() {
                    counter_s[get(&chars, j - 1)] += 1;
                }
            }
        }

        return s[l..r].to_string();
    }

    /// [53. maximum subarray](https://leetcode.cn/problems/maximum-subarray/description)
    ///
    /// Given an integer array nums, find the
    /// subarray
    /// with the largest sum, and return its sum.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
    /// Output: 6
    /// Explanation: The subarray [4,-1,2,1] has the largest sum 6.
    ///
    /// Example 2:
    ///
    /// Input: nums = [1]
    /// Output: 1
    /// Explanation: The subarray [1] has the largest sum 1.
    ///
    /// Example 3:
    ///
    /// Input: nums = [5,4,-1,7,8]
    /// Output: 23
    /// Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     1 <= nums.length <= 10^5
    ///     -10^4 <= nums[i] <= 10^4
    ///
    ///  
    ///
    /// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        if nums.len() == 0 {
            return 0;
        }
        // 状态转移方程:
        // f(i) = max(f(i-1), nums[i])
        //
        // 初始状态:
        // f(0) = nums[0]

        // dp dp 表表示 f(i-1) 的值
        let mut dp = nums[0];
        let mut res = dp;
        for i in 1..nums.len() {
            dp = std::cmp::max(nums[i], dp + nums[i]);
            res = std::cmp::max(res, dp);
        }
        res
    }

    /// [56. Merge Intervals](https://leetcode.cn/problems/merge-intervals/description)
    ///
    /// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
    ///
    /// Example 1:
    ///
    /// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
    /// Output: [[1,6],[8,10],[15,18]]
    /// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
    ///
    /// Example 2:
    ///
    /// Input: intervals = [[1,4],[4,5]]
    /// Output: [[1,5]]
    /// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
    ///
    /// Constraints:
    ///
    ///	1 <= intervals.length <= 10^4
    ///	intervals[i].length == 2
    ///	0 <= starti <= endi <= 10^4
    pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        if intervals.len() <= 1 {
            return intervals;
        }
        // 排序
        intervals.sort_by(|a, b| a[0].cmp(&b[0]));

        // 合并
        let mut res = vec![intervals[0].clone()];
        for i in 1..intervals.len() {
            let last = res.last_mut().unwrap();
            if intervals[i][0] > last[1] {
                res.push(intervals[i].clone());
                continue;
            }
            last[1] = std::cmp::max(last[1], intervals[i][1]);
        }
        return res;
    }

    /// [189. Rotate Array](https://leetcode.cn/problems/rotate-array)
    ///
    /// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
    ///
    /// Example 1:
    ///
    /// Input: nums = [1,2,3,4,5,6,7], k = 3
    /// Output: [5,6,7,1,2,3,4]
    /// Explanation:
    /// rotate 1 steps to the right: [7,1,2,3,4,5,6]
    /// rotate 2 steps to the right: [6,7,1,2,3,4,5]
    /// rotate 3 steps to the right: [5,6,7,1,2,3,4]
    ///
    /// Example 2:
    ///
    /// Input: nums = [-1,-100,3,99], k = 2
    /// Output: [3,99,-1,-100]
    /// Explanation:
    /// rotate 1 steps to the right: [99,-1,-100,3]
    /// rotate 2 steps to the right: [3,99,-1,-100]
    ///
    /// Constraints:
    ///
    ///	1 <= nums.length <= 10^5
    ///	-2^31 <= nums[i] <= 2^31 - 1
    ///	0 <= k <= 10^5
    ///
    /// Follow up:
    ///
    ///	Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
    ///	Could you do it in-place with O(1) extra space?
    pub fn rotate_array(nums: &mut Vec<i32>, k: i32) {
        let k = k as usize;
        let count = Self::gcd(nums.len(), k);
        let k = k % nums.len();

        for i in 0..count {
            let mut tmp = nums[i];
            let mut j = (i + k) % nums.len();
            let mut ok = true;
            while ok {
                (nums[j], tmp) = (tmp, nums[j]);
                ok = j != i;
                j = (j + k) % nums.len();
            }
        }
    }

    /// Greate Common Divisor
    /// 利用欧几里得算法, 计算 a 与 b 的最大公约数
    pub fn gcd(mut a: usize, mut b: usize) -> usize {
        while a != 0 {
            (a, b) = (b % a, a);
        }
        return b;
    }

    /// [238. Product of Array Except Self](https://leetcode.cn/problems/product-of-array-except-self/description)
    ///
    /// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
    ///
    /// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
    ///
    /// You must write an algorithm that runs in O(n) time and without using the division operation.
    ///
    /// Example 1:
    ///
    /// Input: nums = [1,2,3,4]
    /// Output: [24,12,8,6]
    ///
    /// Example 2:
    ///
    /// Input: nums = [-1,1,0,-3,3]
    /// Output: [0,0,9,0,0]
    ///
    /// Constraints:
    ///
    ///	2 <= nums.length <= 105
    ///	-30 <= nums[i] <= 30
    ///	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
    ///
    /// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut left = vec![1; nums.len()];
        let mut right = vec![1; nums.len()];
        for i in 1..nums.len() {
            left[i] = left[i - 1] * nums[i - 1];
            right[nums.len() - 1 - i] = right[nums.len() - i] * nums[nums.len() - i];
        }

        let mut result = Vec::with_capacity(nums.len());
        for i in 0..nums.len() {
            result.push(left[i] * right[i]);
        }
        result
    }

    /// [41. Frist Missing Positive](https://leetcode.cn/problems/first-missing-positive)
    ///
    /// Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.
    ///
    /// You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.
    ///
    /// Example 1:
    ///
    /// Input: nums = [1,2,0]
    /// Output: 3
    /// Explanation: The numbers in the range [1,2] are all in the array.
    ///
    /// Example 2:
    ///
    /// Input: nums = [3,4,-1,1]
    /// Output: 2
    /// Explanation: 1 is in the array but 2 is missing.
    ///
    /// Example 3:
    ///
    /// Input: nums = [7,8,9,11,12]
    /// Output: 1
    /// Explanation: The smallest positive integer 1 is missing.
    ///
    /// Constraints:
    ///
    ///	1 <= nums.length <= 10^5
    ///	-2^31 <= nums[i] <= 2^31 - 1
    pub fn first_missing_positive(mut nums: Vec<i32>) -> i32 {
        for i in 0..nums.len() {
            if nums[i] == 0 {
                nums[i] = -1;
            }
        }

        for i in 0..nums.len() {
            while nums[i] >= 1 && nums[i] <= nums.len() as i32 {
                let j = nums[i] as usize - 1;
                let tmp = nums[j];
                nums[j] = 0;
                if j == i {
                    break;
                }
                if tmp != 0 {
                    nums[i] = tmp;
                } else {
                    nums[i] = -1;
                }
            }
        }

        for (i, &num) in nums.iter().enumerate() {
            if num != 0 {
                return i as i32 + 1;
            }
        }
        nums.len() as i32 + 1
    }

    /// [73. Set Matrix Zeros](https://leetcode.cn/problems/set-matrix-zeroes/description)
    ///
    /// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
    ///
    /// You must do it in place.
    ///
    /// Example 1:
    ///
    /// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
    /// Output: [[1,0,1],[0,0,0],[1,0,1]]
    ///
    /// Example 2:
    ///
    /// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
    /// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
    ///
    /// Constraints:
    ///
    ///	m == matrix.length
    ///	n == matrix[0].length
    ///	1 <= m, n <= 200
    ///	-231 <= matrix[i][j] <= 231 - 1
    ///
    /// Follow up:
    ///
    ///	A straightforward solution using O(mn) space is probably a bad idea.
    ///	A simple improvement uses O(m + n) space, but still not the best solution.
    ///	Could you devise a constant space solution?
    pub fn set_zeros(matrix: &mut Vec<Vec<i32>>) {
        // 标记第一行/第一类是否包含0
        let (mut row0, mut column0) = (false, false);
        for i in 0..matrix.len() {
            for j in 0..matrix[0].len() {
                if matrix[i][j] == 0 {
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                    if i == 0 {
                        row0 = true;
                    }
                    if j == 0 {
                        column0 = true
                    }
                }
            }
        }

        // 遍历赋值
        for i in 1..matrix.len() {
            for j in 1..matrix[i].len() {
                if matrix[i][0] == 0 || matrix[0][j] == 0 {
                    matrix[i][j] = 0;
                }
            }
        }
        if row0 {
            for i in 0..matrix[0].len() {
                matrix[0][i] = 0;
            }
        }
        if column0 {
            for i in 0..matrix.len() {
                matrix[i][0] = 0;
            }
        }
    }

    /// [54. Spiral Matrix](https://leetcode.cn/problems/spiral-matrix/description/)
    ///
    /// Given an m x n matrix, return all elements of the matrix in spiral order.
    ///
    /// Example 1:
    ///
    /// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
    /// Output: [1,2,3,6,9,8,7,4,5]
    ///
    /// Example 2:
    ///
    /// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
    /// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
    ///
    /// Constraints:
    ///
    ///	m == matrix.length
    ///	n == matrix[i].length
    ///	1 <= m, n <= 10
    ///	-100 <= matrix[i][j] <= 100
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        if matrix.is_empty() {
            return vec![];
        }
        let m = matrix.len();
        let n = matrix[0].len();
        let mut res = Vec::with_capacity(m * n);

        let (mut left, mut right, mut top, mut bottom) = (0, n, 0, m);
        while left < right && top < bottom && res.capacity() > res.len() {
            // 向右
            for i in left..right {
                if res.capacity() <= res.len() {
                    break;
                }
                res.push(matrix[top][i]);
            }
            top += 1;
            // 向下
            for i in top..bottom {
                if res.capacity() <= res.len() {
                    break;
                }
                res.push(matrix[i][right - 1]);
            }
            right -= 1;
            // 向左
            for i in (left..right).rev() {
                if res.capacity() <= res.len() {
                    break;
                }
                res.push(matrix[bottom - 1][i]);
            }
            bottom -= 1;
            // 向上
            for i in (top..bottom).rev() {
                if res.capacity() <= res.len() {
                    break;
                }
                res.push(matrix[i][left]);
            }
            left += 1;
        }
        res
    }

    /// [48. Rotate Image](https://leetcode.cn/problems/rotate-image/description)
    ///
    /// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
    ///
    /// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
    ///
    /// Example 1:
    ///
    /// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
    /// Output: [[7,4,1],[8,5,2],[9,6,3]]
    ///
    /// Example 2:
    ///
    /// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
    /// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
    ///
    /// Constraints:
    ///
    ///	n == matrix.length == matrix[i].length
    ///	1 <= n <= 20
    ///	-1000 <= matrix[i][j] <= 1000
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        if matrix.is_empty() {
            return;
        }

        let n = matrix.len();
        for l in 0..(n / 2 + n % 2) {
            for r in (l + 1..n - l).rev() {
                let (mut nr, mut nl) = (l, n - 1 - r);
                let mut tmp = matrix[nr][nl];
                matrix[nr][nl] = matrix[r][l];
                loop {
                    (nr, nl) = (nl, n - 1 - nr);
                    (tmp, matrix[nr][nl]) = (matrix[nr][nl], tmp);
                    if nr == r && nl == l {
                        break;
                    }
                }
            }
        }
    }
}
