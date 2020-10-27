import json

NEW_BOOK_COST = 6
OLD_BOOK_COST = 3.5

with open("booksBenchmark/PythonVersion/coupons.json") as coupons_file:
    COUPONS = json.load(coupons_file)


def compute_cost(new: int, old: int) -> float:
    """Returns the cost of buying `new` new books and `old` old books.
  
    >>> compute_cost(1, 0)
    6
    >>> compute_cost(0, 1)
    3.5
    >>> compute_cost(1, 1)
    9.5
    >>> compute_cost(3, 2)
    25
  """
    return new * NEW_BOOK_COST + old * OLD_BOOK_COST


def is_coupon_valid(coupon_code: str) -> bool:
    """Return True if the provided coupon is valid.
  
    >>> is_coupon_valid('free-stuff')
    True
    >>> is_coupon_valid('half-off')
    True
    >>> is_coupon_valid('anything-else')
    False
    """
    for i in COUPONS:
        if coupon_code == i["Coupon"]:
            return True
    return False


def apply_coupon_discount(cost: float, coupon_code: str) -> float:
    """Return the cost after applying the discount for coupon to cost.
  
    >>> apply_coupon_discount(50.0, 'half-off')
    25.0
    >>> apply_coupon_discount(10.0, 'free-stuff')
    0
    >>> apply_coupon_discount(30.0, 'anything-else')
    30.0
    """
    for i in COUPONS:
        if coupon_code == i["Coupon"]:
            return cost * i["Discount"]
    return cost

