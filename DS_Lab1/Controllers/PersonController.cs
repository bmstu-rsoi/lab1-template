using DS_Lab1.Models;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace DS_Lab1.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class PersonController : Controller
    {
        [HttpGet]
        public Person GetPerson(int personId)
        {

        }

        [HttpGet]
        public IEnumerable<Person> GetPersons()
        {

        }

        [HttpPost]
        public void CreatePerson(Person person)
        {

        }

        [HttpPatch]
        public void UpdatePerson(Person person)
        {

        }

        [HttpDelete]
        public void DeletePerson(int personId) 
        {

        }
    }
}
