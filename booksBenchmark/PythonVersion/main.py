import business_logic
import json


def input_number_of_books(prompt: str) -> int:
    while True:
        response = input(prompt)
        if response.isdigit():
            response = int(response)
            if response >= 0:
                return response
        print("Please select a value of 0 or greater")


def input_has_coupon() -> bool:
    while True:
        has_coupon = input("Do you have a coupon? [Y/N] ").upper()
        if has_coupon in ["N", "Y"]:
            return has_coupon == "Y"
        else:
            print("Please provide valid input.")

with open("booksBenchmark/PythonVersion/purchases.json") as file:
    reader = json.load(file)
    purchases = list(reader)

print("Welcome to Bargain Books")
print(f"New books are ${business_logic.NEW_BOOK_COST:.2f}")
print(f"Old books are ${business_logic.OLD_BOOK_COST:.2f}")
new = input_number_of_books("How many new books are you buying today? ")
old = input_number_of_books("How many old books are you buying today? ")
total = business_logic.compute_cost(new, old)

has_coupon = input_has_coupon()

if has_coupon:
    coupon = input("What is your coupon code? ")
    if business_logic.is_coupon_valid(coupon):
        total = business_logic.apply_coupon_discount(total, coupon)
        print("Coupon successfully applied!")
    else:
        print("I'm sorry. That coupon code is not valid.")
        coupon = None
else:
    coupon = None

print(f"That will be ${total:.2f}")

purchase = {"Old Books": old, "New Books": new, "Coupon": coupon, "Total": total,}
purchases.append(purchase)

with open("booksBenchmark/PythonVersion/purchases.json", "w") as purchases_file:
    json.dump(purchases, purchases_file)
